#!/bin/bash

set -x 

USER="ztrust_rw"
PASS="tA2HCJr1uoN500XWbFAP"
HOST="service-ck-kauo4f9b7a.ck-kauo4f9b7a-hb.jvessel2.jdcloud.com"
DB="ztrust"
MSG="$(date +%s)"

# ./clickhouse-test -host service-ck-kauo4f9b7a.ck-kauo4f9b7a-hb.jvessel2.jdcloud.com -database ztrust -msg 1688552554 -password tA2HCJr1uoN500XWbFAP -username ztrust_rw

./clickhouse-test -host $HOST  -database $DB -msg $MSG -password $PASS -username $USER

