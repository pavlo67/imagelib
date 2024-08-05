Що треба зробити:

    1. Добитись, щоб камеру можна було переглянути за допомогою GStreamer — саме цьому присвячено цей опис.
    2. Добитись, щоб OpenCV було збілджено з підтримкою GStreamer (так, зазвичай, і є, якщо GStreamer 
    ставиться перед OpenCV). Це перевіряється або в процесі білда OpenCV, або спеціяльною програмою, яка описана 
    ось тут: https://learnopencv.com/get-opencv-build-information-getbuildinformation/
    
    Коли виконані ці два пункти, наша програма буде працювати. Инакше — ні (альтернативні варінти — замість 
    GStreamer аналогічно пройти з FFmpeg чи т.п., але ті шляхи я сам ще ні разу не пройшов).


## Початкова перевірка:

    $ rpicam-hello
    $ rpicam-vid   -t 0       
    $ rpicam-still -t 0  

Якщо ці утиліти недоступні, то:

    $ sudo apt install rpicam-apps

Якщо камера працездатна, то rpicam-hello створює вікно з картинкою з неї.


## Драйвери 

За потреби, вони вписуються в /boot/firmware/config.txt. Наприклад, для камери з сенсором imx519 потрібно наступне: 

    [all]
    dtoverlay=imx519


## Бібліотеки

Потрібна libcamera. Її можна перевірити:

    $ libcamera--hello
    $ gst-launch-1.0 libcamerasrc ! queue ! glimagesink

Другий виклик, скоріш за все, не працює по дефолту. Якщо "libcamerasrc: element not found", то:

    $ sudo apt install gstreamer*

Але якщо виклик проходить (наприклад, після "sudo apt install gstreamer*") і генерує помилку: 
"ERROR IPCPipe ipc_pipe_unixsocket.cpp:134 Call timeout!" — то слід збілдити libcamera з сирців. Ця 
помилка розібрана ось тут: https://github.com/raspberrypi/libcamera/issues/115

??? Неясно, чи ці пакети потрібні: 

    $ sudo apt-get install cmake meson flex bison
    $ sudo apt-get install libglib2.0-dev libjpeg-dev libx264-dev
    $ sudo apt-get install libgtk2.0-dev libcanberra-gtk* libgtk-3-dev

!!! Але ясно, що для libcamera необхідні оці два:

    $ sudo apt install libgnutls28-dev
    $ sudo apt install openssl    

Наступна установка Libcamera: https://libcamera.org/getting-started.html

    $ git clone https://git.libcamera.org/libcamera/libcamera.git
    $ cd libcamera
    $ meson setup build
    $ ninja -C build install

Після білду Libcamera її каталог слід залишити на диску, наприклад: /home/pi/0/_install/libcamera/build/

Тоді додати в ~/.bashrc:

    export GST_REGISTRY=/home/pi/0/_install/libcamera/build/src/gstreamer/registry.data
    export GST_PLUGIN_PATH=/home/pi/0/_install/libcamera/build/src/gstreamer

Застосувати без перезавантаження і перевірити:

    $ source ~/.profile
    $ printenv | grep GST
    GST_REGISTRY=/home/pi/0/_install/libcamera/build/src/gstreamer/registry.data
    GST_PLUGIN_PATH=/home/pi/0/_install/libcamera/build/src/gstreamer

Камера нарешті доступна через GStreamer:

    $ gst-launch-1.0 libcamerasrc ! queue ! glimagesink

## Install Golang

Download https://go.dev/dl/go1.22.5.linux-arm64.tar.gz

    $ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.5.linux-arm64.tar.gz
    $ export PATH=$PATH:/usr/local/go/bin
    $ go version
    go version go1.22.5 linux/arm64

Add to ~/.profile

    export PATH=$PATH:/usr/local/go/bin


## Install GoCV/OpenCV

    $ cd ~/go/pkg   

    git clone https://github.com/hybridgroup/gocv.git
    make install_raspi
   

