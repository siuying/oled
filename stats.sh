host=`hostname -I | cut -d' ' -f1`
cpu=`top -bn1 | grep load | awk '{printf "CPU Load: %.2f", $(NF-2)}'`
memory=`free -m | awk 'NR==2{printf "Mem: %s/%sMB %.2f%%", $3,$2,$3*100/$2 }'`
disk=`df -h | awk '$NF=="/"{printf "Disk: %.1f/%.1fGB %s", $3,$2,$5}'`

./oled "./start.ttf" "IP: ${host}" "$cpu" "$memory" "$disk"