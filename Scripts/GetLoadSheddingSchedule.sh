#!/bin/bash

#Check the current user

#The purpose of this script is to automate the process of getting my loadshedding schedule to the NightCiytHomeLab.
CurrentUID="$(id -u)"
if [[ "${CurrentUID}" -eq 0 ]]
then
  echo 'This script must not be run as a root user. Please switch a non-root user.'
fi

MyLocation="${1}"
ESPToken="${2}"

#Check API Allowance 
GetMyBalance="$(curl --request GET 'https://developer.sepush.co.za/business/2.0/api_allowance' --header 'token:'${ESPToken} | tee APIBalance.txt)"
LocalBalance="$(cut -b 23-24 APIBalance.txt | tr --d ,)"
if [[ "${LocalBalance}" -gt 10 ]]
then
   echo 'Current API Call Balance:' $LocalBalance
elif [[ "${LocalBalance}" -lt 10 ]]
then 
    echo 'Current API Call Balance:' $LocalBalance
fi

#Get My LoadShedding Schedule
if [[ "${LocalBalance}" -le 10 ]]
then
    echo 'API Balance is too low or API Balance has been exceeded. Use existing outputs instead!'
elif [[ "${LocalBalance}" -gt 10 ]]
then
    GetCurrentSchedule="$(curl --location --request GET 'https://developer.sepush.co.za/business/2.0/area?id='${MyLocation}'&test=current' --header 'token:'${ESPToken} | tee CurrentLoadSheddingSchedule.txt)"
    GetFutureSchedule="$(curl --location --request GET 'https://developer.sepush.co.za/business/2.0/area?id='${MyLocation}'&test=future'  --header 'token:'${ESPToken} | tee UpcomingLoadSheddingSchedule.txt)"
    echo 'API call made successfully'
fi


if [[ "${?}" -eq 0 ]]
then
    echo 'Script has been executed successfully.'
elif [[ "${?}" -ge 1 ]]
then
    echo 'Script has not been executed successfully.'
    exit 1
fi

