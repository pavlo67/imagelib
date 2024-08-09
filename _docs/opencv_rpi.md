## Загальна послідовність

Зачистити дефолтну інсталяцію:

    $ sudo apt remove libopencv*

Перевірити, що чисто (перевірки мають показувати "нічого"):

    $ dpkg -l libopencv-dev
    $ sudo find / -name "libopencv_core.so.*"
    $ sudo find / -name "libopencv*"          # may show docs only

Поставити Golang.

    $ cd ~/Downloads

Завантажити архів: https://go.dev/dl/go1.22.5.linux-arm64.tar.gz

    $ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.5.linux-arm64.tar.gz
    $ export PATH=$PATH:/usr/local/go/bin
    $ go version

Очікуваний результат:

    go version go1.22.5 linux/arm64

Додати в ~/.profile: 

    export PATH=$PATH:/usr/local/go/bin

Поставити GoCV + OpenCV (разом). Для цього спершу слід створити каталог ~/go/pkg/mod/gocv.io/x (послідовно створюючи відповідні над-каталоги).    

    $ cd ~/go/pkg/mod/gocv.io/x
    $ git clone https://github.com/hybridgroup/gocv.git
    $ cd gocv
    $ make install_raspi

Очікуваний результат:

    gocv version: 0.37.0
    opencv lib version: 4.10.0


