import re
import ssl
import nltk
from nltk.corpus import words

try:
    _create_unverified_https_context = ssl._create_unverified_context
except AttributeError:
    pass
else:
    ssl._create_default_https_context = _create_unverified_https_context

nltk.download("words")

def check_spelling(content):
    """Check spelling of words in the given content.

    Args:
        content (str): The text to check.

    Returns:
        list: A list of words with spelling errors.
    """
    errors = []
    content = content.strip()

    words_in_text = content.split(" ")
    for word in words_in_text:
        clean_word = re.sub(r"[^\w]", "", word.lower())
        if clean_word and clean_word not in words.words():
            errors.append(word)
    
    return errors

# Example usage in Flask
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/check_spelling', methods=['POST'])
def spelling_check_endpoint():
    data = request.json
    content = data.get("text", "")
    errors = check_spelling(content)
    return jsonify({"errors": errors})

if __name__ == "__main__":
    app.run(debug=True)
