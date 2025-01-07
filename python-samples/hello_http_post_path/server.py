from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/hello/<name>', methods=['GET'])
def hello(name):
    if not name:
        return "Hello, World (no Name parm)!", 200
    return f"Hello, {name}!", 200

if __name__ == '__main__':
    app.run(port=8080)
