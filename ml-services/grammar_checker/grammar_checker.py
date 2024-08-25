import language_tool_python
import tkinter as tk
from tkinter.scrolledtext import ScrolledText

class GrammarChecker:

    def __init__(self):
        self.tool = language_tool_python.LanguageTool('en-US')
        self.root = tk.Tk()
        self.root.title("Grammar Checker")
        self.root.geometry("600x500")

        self.text = ScrolledText(self.root, font=("Arial", 14), wrap=tk.WORD)
        self.text.bind("<KeyRelease>", self.check_grammar)
        self.text.pack(expand=True, fill=tk.BOTH)
        self.old_text = ""

        self.root.mainloop()

    def check_grammar(self, event):
        content = self.text.get("1.0", tk.END).strip()


        if content != self.old_text:
            self.old_text = content


            for tag in self.text.tag_names():
                self.text.tag_delete(tag)

           
            matches = self.tool.check(content)
            for match in matches:
                start_idx = self.text.index(f"1.0 + {match.offset} chars")
                end_idx = self.text.index(f"1.0 + {match.offset + match.errorLength} chars")
                self.text.tag_add(str(match), start_idx, end_idx)
                self.text.tag_config(str(match), foreground="blue")

if __name__ == "__main__":
    GrammarChecker()
