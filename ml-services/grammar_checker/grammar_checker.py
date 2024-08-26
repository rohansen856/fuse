# import language_tool_python
# import tkinter as tk
# from tkinter.scrolledtext import ScrolledText

# class GrammarChecker:

#     def __init__(self):
#         self.tool = language_tool_python.LanguageTool('en-US')
#         self.root = tk.Tk()
#         self.root.title("Grammar Checker")
#         self.root.geometry("600x500")

#         self.text = ScrolledText(self.root, font=("Arial", 14), wrap=tk.WORD)
#         self.text.bind("<KeyRelease>", self.check_grammar)
#         self.text.pack(expand=True, fill=tk.BOTH)
#         self.old_text = ""

#         self.root.mainloop()

#     def check_grammar(self, event):
#         content = self.text.get("1.0", tk.END).strip()


#         if content != self.old_text:
#             self.old_text = content


#             for tag in self.text.tag_names():
#                 self.text.tag_delete(tag)

           
#             matches = self.tool.check(content)
#             for match in matches:
#                 start_idx = self.text.index(f"1.0 + {match.offset} chars")
#                 end_idx = self.text.index(f"1.0 + {match.offset + match.errorLength} chars")
#                 self.text.tag_add(str(match), start_idx, end_idx)
#                 self.text.tag_config(str(match), foreground="blue")

# if __name__ == "__main__":
#     GrammarChecker()

from flask import Flask, request, jsonify
import re
import nltk
from nltk.corpus import words
import ssl

# Fix SSL certificate issue for NLTK downloads
try:
    _create_unverified_https_context = ssl._create_unverified_context
except AttributeError:
    pass
else:
    ssl._create_default_https_context = _create_unverified_https_context

nltk.download("words")

app = Flask(__name__)

# POST route for spelling checking
@app.route('/spelling', methods=['POST'])
def check_spelling():
    data = request.json  # assuming the data is sent as JSON
    text = data.get("text", "")
    words_in_text = text.split(" ")

    errors = []
    for word in words_in_text:
        clean_word = re.sub(r"[^\w]", "", word.lower())
        if clean_word and clean_word not in words.words():
            start_idx = text.find(word)
            errors.append({
                "word": word,
                "start_idx": start_idx,
                "end_idx": start_idx + len(word),
            })

    response = {"errors": errors}
    return jsonify(response)

if __name__ == '__main__':
    app.run(debug=True)