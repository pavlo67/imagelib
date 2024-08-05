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

## GoCV / RPI

Опис: https://pkg.go.dev/gocv.io/x/gocv#readme-raspbian
    
    $ cd $GOPATH/pkg/mod/gocv.io/x/
    $ git clone https://github.com/hybridgroup/gocv.git
    $ cd gocv
    $ make install_raspi

## OpenCV from source / Raspberry Pi 5

Ручний білд з сирців (про инші способи — геть всі неробочі — в наступному пункті) згідно з https://pypi.org/project/opencv-contrib-python/:

    # слід зачистити попередні інсталяції OpenCV, як описано вище
    # слід інсталювати-перевірити gstreamer    
    
    $ git clone --recursive https://github.com/opencv/opencv-python.git
    $ cd opencv-python
    $ export ENABLE_CONTRIB=1
    $ pip wheel . --verbose

    # після того, як відпрацював cmake і почався білд — рядки в консолі на кшталт "[11/1389] Building ..." 
    # варто відмотати консоль нагору і перевірити настройки білда — "GSTREAMER=ON" і все, що стосується python — повинна бути вказана саме та версія, з якою ми працюємо.

    $ pip install *.whl

При цьому openCV-бібліотеки генеруються як .a-файли в build-підкаталогах opencv-python — напевне, вони не будуть працювати ні з чим, окрім python'а. І, ймовірно, каталог opencv-python не можна прибирати з системи.

## OpenCV from source / common

Manual: https://docs.opencv.org/4.x/d7/d9f/tutorial_linux_install.html

    $ sudo make install
    # sudo make uninstall 

# GStreamer support

Check gstreamer support (python code)

    import cv2
    print(cv2.getBuildInformation())

manuals

    https://galaktyk.medium.com/how-to-build-opencv-with-gstreamer-b11668fa09c
    https://medium.com/@arfanmahmud47/build-opencv-4-from-source-with-gstreamer-ubuntu-zorin-peppermint-c2cff5393ef

# OpenCV from apt repos:

!!! Але в цій версії OpenCV приходить застаріле 

    sudo apt install libopencv-dev
    sudo apt install python3-opencv
    sudo apt install python3-opencv-contrib # ??? python3-contrib-opencv


## OpenCV for Python only from pip repos:

!!! Але в цій версії OpenCV приходить без підтримки GStreamer

    pip3 install opencv-contrib-python

## Raspberry Pi 5 — невдалі спроби

На жаль, жоден з трьох способів, описаних тут —  https://qengineering.eu/install%20opencv%20on%20raspberry%20pi%205.html — не спрацював для нашої конфігурації.   

У нас на RPI5 два пітони:
* системний 3.11, прихований за допомогою pyenv;
* додатковий 3.9.16, поставлений з pyenv як дефолтний, він необхідний для роботи з pycoral.

Отже:
* apt-пакет — ставиться, але в /usr/lib/python — для 3.11, це не наше;
* pip-пакет — ставиться для нашого пітона і запускається, але він збілджений без підтримки gstreamer — нашу задачу він не виконає :-(
* білд openCV з сирців чомусь не генерує артефакти для пітону (відсутній cv2...so-файл бібліотеки, але вдавалось згенерити опис в site-packages — з ним "include cv2" спрацьовує, але самі функції з cv2 не викликаються).

В останньому випадку окрім початкової конфігурації було проведено багато спроб із різними cmake-ключами. Правдоподібно виглядає наступний опис:

    -D WITH_PYTHON=ON \
    -D BUILD_opencv_python3=ON \
    -D BUILD_opencv_python2=OFF \
    -D PYTHON3_EXECUTABLE=$(which python) \
    -D OPENCV_PYTHON3_INSTALL_PATH=$(python -c "from distutils.sysconfig import get_python_lib; print(get_python_lib())") \
    -D PYTHON3_INCLUDE_DIR=$(python -c "from distutils.sysconfig import get_python_inc; print(get_python_inc())") \
    -D PYTHON3_PACKAGES_PATH=$(python -c "from distutils.sysconfig import get_python_lib; print(get_python_lib())") \
    
— але він теж відпрацював без істотних помилок в консолі не згенеривши при цьому потрібний результат :-(

## opencv-python / cv2 docs

На жаль, адекватної документації в природі не існує. Ось дискусія щодо цього: https://github.com/opencv/opencv-python/issues/522


