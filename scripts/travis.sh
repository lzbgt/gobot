#!/bin/bash
PACKAGES=('gobot' 'gobot/api' 'gobot/platforms/intel-iot/edison' 'gobot/sysfs' $(ls ./platforms | sed -e 's/^/gobot\/platforms\//'))
EXITCODE=0

if [ "$(go version | cut -d' ' -f3)" = "devel" ]
then
  go get golang.org/x/tools/cmd/cover
else
  go get code.google.com/p/go.tools/cmd/cover
fi

echo "mode: count" > profile.cov
touch tmp.cov
for package in "${PACKAGES[@]}"
do
  go test -covermode=count -coverprofile=tmp.cov github.com/hybridgroup/$package
  if [ $? -ne 0 ]
  then
    EXITCODE=1
  fi
  cat tmp.cov | grep -v "mode: count" >> profile.cov
done

if [ $EXITCODE -ne 0 ]
then
  exit $EXITCODE
fi

export PATH=$PATH:$HOME/gopath/bin/
goveralls -coverprofile=profile.cov -service=travis-ci -repotoken=sFrR9ZmLP5FLc34lOaqir67RPzYOvFPUB
