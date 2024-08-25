from flask import Flask, request, jsonify

app = Flask(__name__)

# GET route
@app.route('/grammar', methods=['GET'])
def get_data():
    # You can fetch and return some data here
    data = {"message": "This is a GET request"}
    return jsonify(data)

# POST route
@app.route('/grammar', methods=['POST'])
def post_data():
    # Retrieve data from the POST request
    data = request.json  # assuming the data is sent as JSON
    response = {"received_data": data}
    return jsonify(response)

if __name__ == '__main__':
    app.run(debug=True)
