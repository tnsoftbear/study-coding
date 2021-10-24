# Read me (rus)

 #cookie #samesite #php #lax #strict #secure

Это адаптированный на языке PHP пример использования разных значений для параметра SameSite куки. Он основывается на примере [из этого видео](https://github.com/hnasr/javascript_playground/blob/master/samesite/index.js).  
Оригинальный код на джс [смотреть здесь](https://github.com/hnasr/javascript_playground/tree/master/samesite)  

Создал 2 виртуальных хоста в apache: samesite.lv и othersite.lv для https протокола. Потому что файлы cookie с `SameSite=None` также должны быть указаны `Secure`, что означает, что они требуют безопасного контекста.

```conf
<VirtualHost *:443>
      DocumentRoot "E:\\web\\tn_samples\\cookie\\samesite\\"
      ServerName samesite.lv
      SSLCertificateFile "E:/dev/certs/samesite.lv.crt"
      SSLCertificateKeyFile "E:/dev/certs/samesite.lv.key"
</VirtualHost>

<VirtualHost *:443>
      DocumentRoot "E:\\web\\tn_samples\\cookie\\samesite\\othersite\\"
      ServerName othersite.lv
      SSLCertificateFile "E:/dev/certs/samesite.lv.crt"
      SSLCertificateKeyFile "E:/dev/certs/samesite.lv.key"
</VirtualHost>
```

Создал ssl ключи

```sh
E:/dev/apache/bin/openssl.exe req -x509 -newkey rsa:4096 -sha256 -nodes -keyout samesite.lv.key -out samesite.lv.crt -days 3650 -config E:/dev/apache/conf/openssl.cnf
```

Для включения ssl в apache в `httpd.conf` включил модули `mod_socache_shmcb.so`, `mod_ssl.so` и сделал инклуд `httpd-ssl.conf`, из которого удалил весь дефалтный `<VirtualHost *:443>...</VirtualHost>`

```conf
LoadModule socache_shmcb_module modules/mod_socache_shmcb.so
LoadModule ssl_module modules/mod_ssl.so
Include conf/extra/httpd-ssl.conf 
```

## Links

* [php: setcookie()](https://www.php.net/manual/en/function.setcookie.php)
* <https://web.dev/samesite-cookies-explained/>