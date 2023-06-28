from flask import Flask
from prisma import Prisma, register

from routes.post import post_blueprint

db = Prisma()

db.connect()
register(db)

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello world</p>"

app.register_blueprint(post_blueprint, url_prefix='/post')


if __name__ == "main":
    app.run(debug=True, port= 5000, threaded=True)