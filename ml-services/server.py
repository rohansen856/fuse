from flask import Flask, request, jsonify
import language_tool_python

app = Flask(_name_)
tool = language_tool_python.LanguageTool('en-US')

# GET route
@app.route('/grammar', methods=['GET'])
def get_data():
    data = {"message": "This is a GET request"}
    return jsonify(data)

# POST route for grammar checking
@app.route('/grammar', methods=['POST'])
def post_data():
    data = request.json  # assuming the data is sent as JSON
    text = data.get("text", "")
    matches = tool.check(text)
    errors = [{"message": match.message, "offset": match.offset, "errorLength": match.errorLength} for match in matches]
    response = {"errors": errors}
    return jsonify(response)

if _name_ == '_main_':
    app.run(debug=True)
