# OpenCV

## Common actions

check if installed:

    $ dpkg -l libopencv-dev
    $ sudo find / -name "libopencv_core.so.*"
    $ sudo find / -name "libopencv*"          # may show docs only

purge old:

    $ sudo apt remove libopencv*
    $ sudo find / -name "libopencv*"          # may show docs only
    $ sudo rm <path>/libopencv*               # for each path containing "libopencv*" found before


## GoCV

    $ go get -u -d gocv.io/x/gocv
    $ cd $GOPATH/pkg/mod/gocv.io/x/gocv@...
    # vi Makefile
    distro_deps=deps_ubuntu_jammy
    $ make install


## OpenCV from source / common

Manual: https://docs.opencv.org/4.x/d7/d9f/tutorial_linux_install.html

    $ sudo make install
    # sudo make uninstall 

## GStreamer support

Check gstreamer support (python code)

    import cv2
    print(cv2.getBuildInformation())

Manuals

    https://galaktyk.medium.com/how-to-build-opencv-with-gstreamer-b11668fa09c
    https://medium.com/@arfanmahmud47/build-opencv-4-from-source-with-gstreamer-ubuntu-zorin-peppermint-c2cff5393ef

## OpenCV from apt repos:

!!! Але в цій версії OpenCV приходить застаріле 

    sudo apt install libopencv-dev
    sudo apt install python3-opencv
    sudo apt install python3-opencv-contrib # ??? python3-contrib-opencv

