from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/hello', methods=['POST'])
def hello():
    # Parse JSON request body
    data = request.get_json()
    
    # Check if 'name' is provided in the request body
    name = data.get('name', '')
    
    if not name:
        return jsonify(message="Hello, World (no Name parm)!"), 200
    
    return jsonify(message=f"Hello, {name}!")

if __name__ == '__main__':
    app.run(debug=True)
