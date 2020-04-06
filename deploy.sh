wget http://gosspublic.alicdn.com/ossutil/1.6.10/ossutil64
chmod 755 ossutil64
./ossutil64 config -e oss-cn-beijing.aliyuncs.com -i $ACCESSKEY -k $SECRETKEY
./ossutil64 cp -f ./kan oss://kan-bin/