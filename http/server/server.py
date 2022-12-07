# coding: utf-8

import os
import time
import json

import tornado.web
import tornado.options
import tornado.httpserver
from tornado import ioloop

import logger

log = logger.install('stdout')


class TestHandler(tornado.web.RequestHandler):

    def get(self):
        log.debug('-------- test get --------')
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        ret = {
            'code': '0000',
            'msg': 'ok',
            'err': '',
            'data': [
                {'id': 1, 'name': 'gaojs'},
                {'id': 2, 'name': 'love'},
                {'id': 3, 'name': 'zhangxf'},
            ],
        }
        self.write(json.dumps(ret))
        self.finish()

    def post(self):
        log.debug('-------- test post --------')
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        ret = {
            'code': '0000',
            'msg': 'ok',
            'err': '',
            'data': [
                {'id': 1, 'name': 'gaojs'},
                {'id': 2, 'name': 'love'},
                {'id': 3, 'name': 'zhangxf'},
            ],
        }
        self.write(json.dumps(ret))
        self.finish()

    def put(self):
        log.debug('-------- test put --------')
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        ret = {
            'code': '0000',
            'msg': 'ok',
            'err': '',
            'data': [
                {'id': 1, 'name': 'gaojs'},
                {'id': 2, 'name': 'love'},
                {'id': 3, 'name': 'zhangxf'},
            ],
        }
        self.write(json.dumps(ret))
        self.finish()

    def delete(self):
        log.debug('-------- test delete --------')
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        ret = {
            'code': '0000',
            'msg': 'ok',
            'err': '',
            'data': [
                {'id': 1, 'name': 'gaojs'},
                {'id': 2, 'name': 'love'},
                {'id': 3, 'name': 'zhangxf'},
            ],
        }
        self.write(json.dumps(ret))
        self.finish()


class TestTimeoutHandler(tornado.web.RequestHandler):

    def get(self):
        log.debug('-------- test timeout get --------')
        time.sleep(2)
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        ret = {
            'code': '0000',
            'msg': 'ok',
            'err': '',
            'data': [
                {'id': 1, 'name': 'gaojs'},
                {'id': 2, 'name': 'love'},
                {'id': 3, 'name': 'zhangxf'},
            ],
        }
        self.write(json.dumps(ret))
        self.finish()

    def post(self):
        log.debug('-------- test timeout post --------')
        time.sleep(2)
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        ret = {
            'code': '0000',
            'msg': 'ok',
            'err': '',
            'data': [
                {'id': 1, 'name': 'gaojs'},
                {'id': 2, 'name': 'love'},
                {'id': 3, 'name': 'zhangxf'},
            ],
        }
        self.write(json.dumps(ret))
        self.finish()

    def put(self):
        log.debug('-------- test timeout put --------')
        time.sleep(2)
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        ret = {
            'code': '0000',
            'msg': 'ok',
            'err': '',
            'data': [
                {'id': 1, 'name': 'gaojs'},
                {'id': 2, 'name': 'love'},
                {'id': 3, 'name': 'zhangxf'},
            ],
        }
        self.write(json.dumps(ret))
        self.finish()

    def delete(self):
        log.debug('-------- test timeout delete --------')
        time.sleep(2)
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        ret = {
            'code': '0000',
            'msg': 'ok',
            'err': '',
            'data': [
                {'id': 1, 'name': 'gaojs'},
                {'id': 2, 'name': 'love'},
                {'id': 3, 'name': 'zhangxf'},
            ],
        }
        self.write(json.dumps(ret))
        self.finish()

class TestRedicrctHandler(tornado.web.RequestHandler):

    def get(self):
        log.debug('-------- test get --------')
        log.debug("request :%s", dir(self.request))
        log.debug("request params:%s", self.request.arguments)
        log.debug("request body:%s", self.request.body)
        log.debug("request cookies:%s", self.request.cookies)
        log.debug("request headers:%s", self.request.headers)
        self.redirect('http://127.0.0.1:8080/test?a=1')




urls = [
    ('/test', TestHandler),
    ('/test_timeout', TestTimeoutHandler),
    ('/test_redirect', TestRedicrctHandler),
]

class ApiCenterHttp():
    def app(self):
        app = tornado.web.Application(
            handlers  =urls,
            autoreload=False,
            xheaders=True
        )   
        return app 

#tornado.options.define('logging', default='info')

if __name__ == '__main__':
    #tornado.options.parse_command_line()   # 不打印tornado日志，因为与我的logger模块不兼容
    http_server = tornado.httpserver.HTTPServer(ApiCenterHttp().app())
    http_server.bind(8080)
    http_server.start(1)
    log.info('http server start at %d', 8080)
    ioloop.IOLoop.current().start()



