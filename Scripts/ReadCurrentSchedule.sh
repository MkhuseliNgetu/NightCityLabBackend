#!/bin/bash

#Check the current user

#The purpose of this script is to automate the process of getting my loadshedding schedule to the NightCiytHomeLab.
CurrentUID="$(id -u)"
if [[ "${CurrentUID}" -eq 0 ]]
then
  echo 'This script must not be run as a root user. Please switch a non-root user.'
fi


Counter="$(echo "$StartTimes" | wc -l)"
counter=2

until [[ "${counter}" -eq 0 ]]
do
    if [[ $counter -eq 1 ]]
    then
        StartTimes="$(grep -Po '"start": *\K"[^"]*"' CurrentLoadSheddingSchedule.json | tail -n $counter | cut -b 12-17 | sed 's/T/ /g')"
        EndTimes="$(grep -Po '"end": *\K"[^"]*"' CurrentLoadSheddingSchedule.json | tail -n $counter | cut -b 12-17 | sed 's/T/ /g')"
        Stages="$(grep -Po '"note": *\K"[^"]*"' CurrentLoadSheddingSchedule.json | tail -n $counter)"

        echo "Current Starting Time:" ${StartTimes} 
        echo "Current Ending Time:" ${EndTimes}
        echo "Current Stage:" ${Stages}

    fi

((counter--))
done


if [[ "${?}" -eq 0 ]]
then
    echo 'Script has been executed successfully.'
elif [[ "${?}" -ge 1 ]]
then
    echo 'Script has not been executed successfully.'
    exit 1
fi