#!/usr/bin/env bash

aem_dir="/opt/aem"
crx_quickstart_dir="${aem_dir}/crx-quickstart"
license_path="${aem_dir}/license.properties"

# Check if AEM is installed.
# If AEM is not installed e.g. the crx-quickstart directory does not exist.
echo "Checking if AEM is installed"
if [[ ! -d ${crx_quickstart_dir} || $(find ${crx_quickstart_dir} -maxdepth 1 | wc -l) -eq 0 ]]; then
    echo "AEM is not installed"

    jars=$(find ${aem_dir} -type f -maxdepth 1 -name "*.jar")
    num_jars=$(echo "${jars}" | wc -l)

    # Error if there is more than one jar file.
    if (( num_jars > 1 )); then
        echo "/opt/aem cannot contain more than one jar"
        exit 1
    fi

    # If there are no jars, and crx-quickstart wasn't populated by a client (API, CLI), throw an error.
    if (( num_jars == 0 )); then
        echo "No jar files found"
        exit 1
    fi

    aem_jar=${jars}

    echo "Unpacking AEM jar"
    if ! java -Djava.awt.headless=true -jar "${aem_jar}" -unpack; then
        echo "Failed to unpack AEM jar"
        exit 1
    fi

    # Error if license file does not exist.
    if [[ ! -a "${license_path}" ]]; then
        echo "No license file found"
        exit 1
    fi
else
    echo "AEM is already installed"
fi

# Send logs to stdout and stderr.
tail -n+1 -F ${crx_quickstart_dir}/logs/stdout.log > /proc/1/fd/1 &
tail -n+1 -F ${crx_quickstart_dir}/logs/error.log > /proc/1/fd/2 &

echo "Starting AEM"
$crx_quickstart_dir/bin/start

# Dumb-init might make all this unnecessary 
function quit {
    ./stop_aem.sh
    exit 0
}

trap quit SIGINT SIGTERM
while true; do sleep 1; done
