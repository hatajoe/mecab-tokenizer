# ref: http://qiita.com/milk1000cc/items/80c4238675e396d93381

set -x
set -e

BASE_DIR="$HOME/lib/mecab"

sudo apt-get remove mecab

if [ ! -d "$BASE_DIR" ]
then
    mkdir -p $BASE_DIR
fi

if [ ! -f "$BASE_DIR/mecab-0.996.tar.gz" ]
then
    cd $BASE_DIR
    wget http://mecab.googlecode.com/files/mecab-0.996.tar.gz
    tar zxvf mecab-0.996.tar.gz
    cd mecab-0.996
    ./configure
    make
fi
cd $BASE_DIR/mecab-0.996
sudo make install
sudo sh -c "echo '/usr/local/lib' >> /etc/ld.so.conf"
sudo ldconfig

if [ ! -f "$BASE_DIR/mecab-ipadic-2.7.0-20070801.tar.gz" ]
then
    cd $BASE_DIR
    wget http://mecab.googlecode.com/files/mecab-ipadic-2.7.0-20070801.tar.gz
    tar zxvf mecab-ipadic-2.7.0-20070801.tar.gz
    cd mecab-ipadic-2.7.0-20070801
    ./configure --with-charset=utf8
    make
fi
cd $BASE_DIR/mecab-ipadic-2.7.0-20070801
sudo make install
