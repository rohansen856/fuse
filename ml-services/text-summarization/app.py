import requests

def summarize_text(text, api_token, model="facebook/bart-large-cnn"):
    API_URL = f"https://api-inference.huggingface.co/models/{model}"
    headers = {"Authorization": f"Bearer {api_token}"}

    payload = {
        "inputs": text,
        "parameters": {"min_length": 50, "max_length": 200},
    }

    response = requests.post(API_URL, headers=headers, json=payload)
    summary = response.json()

    return summary[0]['summary_text'] if isinstance(summary, list) else summary

# Replace 'your_api_token' with your actual Hugging Face API token
text = """
The Hugging Face Transformers library provides thousands of pre-trained models to perform tasks on different modalities such as text, vision, and audio. These models can be applied to various downstream tasks such as text classification, information extraction, question answering, summarization, translation, and text generation in over 100 languages.
"""

api_token="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"
summary = summarize_text(text, api_token)
print("Summary:", summary)
