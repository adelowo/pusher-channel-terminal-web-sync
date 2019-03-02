while true
do
    uuid=`od -x /dev/urandom | head -1 | awk '{OFS="-"; print $2$3,$4,$5,$6,$7$8$9}'`
    echo "New UUID generated --- $uuid"
    sleep 1
done
