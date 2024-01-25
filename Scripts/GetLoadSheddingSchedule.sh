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
if [[ "${LocalBalance}" -lt 50 ]]
then
   echo 'Current API Call Balance:' $LocalBalance
elif [[ "${LocalBalance}" -eq 50 ]]
then 
    echo 'Current API Call Balance:' $LocalBalance
fi

#Get My LoadShedding Schedule
if [[ "${LocalBalance}" -lt 50 ]]
then
    GetCurrentSchedule="$(curl --location --request GET 'https://developer.sepush.co.za/business/2.0/area?id='${MyLocation} --header 'token:'${ESPToken} > CurrentLoadSheddingSchedule.json)"
    echo 'API call made successfully'
elif [[ "${LocalBalance}" -eq 50 ]]
then
      echo 'API Balance has been exceeded. Using existing outputs instead!'
fi

if [[ "${?}" -eq 0 ]]
then
    echo 'Script has been executed successfully.'
elif [[ "${?}" -ge 1 ]]
then
    echo 'Script has not been executed successfully.'
    exit 1
fi

