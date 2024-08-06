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


## Raspberry Pi 5: OpenCV+Python from source 

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
