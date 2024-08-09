Що треба зробити:

    1. Добитись, щоб камеру можна було переглянути за допомогою GStreamer — саме цьому присвячено цей опис.
    2. Добитись, щоб OpenCV було збілджено з підтримкою GStreamer (так, зазвичай, і є, якщо GStreamer 
    ставиться перед OpenCV). Це перевіряється або в процесі білда OpenCV, або спеціяльною програмою, яка описана 
    ось тут: https://learnopencv.com/get-opencv-build-information-getbuildinformation/
    
    Коли виконані ці два пункти, то звʼязка "Go/C++ + OpenCV + камера" буде працювати. 

## Початкова перевірка:

    $ rpicam-hello
    $ rpicam-vid   -t 0       
    $ rpicam-still -t 0  

Якщо ці утиліти недоступні, то:

    $ sudo apt install rpicam-apps

Якщо камера працездатна, то rpicam-hello створює вікно з картинкою з неї.


## Драйвери 

За потреби, вони вписуються в /boot/firmware/config.txt. Наприклад, для камери з сенсором imx519 (якщо вона не детектується автоматично) потрібно наступне: 

    [all]
    dtoverlay=imx519


## Бібліотеки

Потрібна libcamera. Її можна перевірити:

    $ libcamera--hello
    $ gst-launch-1.0 libcamerasrc ! queue ! glimagesink

Другий виклик, скоріш за все, не працює по дефолту. Якщо "libcamerasrc: element not found", то:

    $ sudo apt install gstreamer*

Але якщо виклик проходить (наприклад, після "sudo apt install gstreamer*") і генерує помилку: 
"ERROR IPCPipe ipc_pipe_unixsocket.cpp:134 Call timeout!" — то слід збілдити libcamera з сирців 
з попередньою установкою додаткових пакетів. Ця помилка розібрана ось тут: 
https://github.com/raspberrypi/libcamera/issues/115

Щоб в системі не залишались дефолтні пакети libcamera, напевне, варто їх прибрати:

    $ sudo apt remove libcamera*

Тоді встановити додаткові пакети:

    $ sudo apt install cmake meson
    $ sudo apt install libgnutls28-dev
    $ sudo apt install openssl    

І встановити Libcamera з сирців: https://libcamera.org/getting-started.html

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

