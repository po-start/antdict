#! /usr/bin/env bash

section="# ----- [ Govendor init ] ----- #"
echo ${section}
govendor init
section="# ----- [ Govendor fetch echo ] ----- #"
echo ${section}
govendor fetch github.com/labstack/echo
section="# ----- [ Govendor fetch echo/middleware ] ----- #"
echo ${section}
govendor fetch github.com/labstack/echo/middleware
section="# ----- [ Govendor fetch gracehttp ] ----- #"
echo ${section}
govendor fetch github.com/facebookgo/grace/gracehttp
section="# ----- [ Govendor fetch xorm ] ----- #"
echo ${section}
govendor fetch github.com/go-xorm/xorm
section="# ----- [ Govendor fetch elastic ] ----- #"
echo ${section}
govendor fetch gopkg.in/olivere/elastic.v5
section="# ----- [ Govendor fetch yaml.v2 ] ----- #"
echo ${section}
govendor fetch gopkg.in/yaml.v2
section="# ----- [ Govendor fetch golog ] ----- #"
echo ${section}
govendor fetch github.com/siddontang/go/log
section="# ----- [ Govendor fetch go-mysql-driver ] ----- #"
echo ${section}
govendor fetch github.com/go-sql-driver/mysql
