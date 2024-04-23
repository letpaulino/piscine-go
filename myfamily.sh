#!/bin/bash
curl -s https://01.gritlab.ax/assets/superhero/all.json | jq --arg hero_id "$HERO_ID" 'map(select(.id == ($hero_id|tonumber)))' | jq  '.[] | .connections.relatives' | sed 's/"//g'
