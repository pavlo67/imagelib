## GStreamer on RPI

Керівництво: https://qengineering.eu/install-gstreamer-1.18-on-raspberry-pi-4.html

Але простіше:

    $ sudo apt install gstreamer*


## How to avoid bufferization on receive

https://forums.developer.nvidia.com/t/how-to-eliminate-gstreamer-camera-buffer/75608

    appsink max-buffers=1 drop=True

Send/receive example

Потокова передача відео з файлу:

    $ gst-launch-1.0 -v filesrc location=T80_360_1_10fps.mp4 ! decodebin ! x264enc ! rtph264pay ! udpsink host=127.0.0.1 port=5600

Отримання потоку:

    $ gst-launch-1.0 -vc udpsrc port=5600 close-socket=false auto-multicast=true ! application/x-rtp, payload=96 ! rtph264depay ! decodebin3 ! fpsdisplaysink sync=false

