#!/usr/bin/env bash

./aem/crx-quickstart/bin/stop

sleep 1
echo -n "Waiting for AEM to stop..."
while true; do
    pid=$(ps -ax | grep java | grep -v "grep" | awk '{print $1}')
    if [[ -z "$pid" ]]; then
        break
    fi
    echo -n "."
    sleep 1
done
echo -e "\nAEM stopped!"
