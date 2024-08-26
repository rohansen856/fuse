import re
from flask import Flask, request, jsonify
import language_tool_python
from nltk.corpus import words
# from .text_summarizer import textsummarization
# from .News_Recommendation import News_Recommendation_Final

app = Flask(__name__)
tool = language_tool_python.LanguageTool('en-US')

# GET route
@app.route('/summarize', methods=['POST'])
def get_data():
    data = request.json
    text = data.get("text", "")
        # -*- coding: utf-8 -*-
    """TextSummarization(NLP).ipynb

    Automatically generated by Colab.

    Original file is located at
        https://colab.research.google.com/drive/1ycFPExm-v7EMDfaw3DDo79RBp5-7qjk-
    """

    import numpy as np
    import pandas as pd
    import warnings
    import re
    import nltk
    from nltk import word_tokenize
    # splits given text into words
    from nltk.tokenize import sent_tokenize
    # splits given text into sentences
    from textblob import TextBlob
    import string
    from string import punctuation
    # This is a pre defined string that contains all the punctuation
    from nltk.corpus import stopwords

    from statistics import mean
    from heapq import nlargest
    # This is used to find n largest elements from a list or something ,
    # it internally uses heap data structure
    from wordcloud import WordCloud
    # These are visual representations of text in a document
    import seaborn as sns

    import matplotlib.pyplot as plt

    import nltk
    nltk.download('stopwords')
    nltk.download('punkt')
    nltk.download('punkt_tab')

    stop_words=set(stopwords.words('english'))
    punctuation = punctuation + '\n' + '—' + '“' + ',' + '”' + '‘' + '-' + '’'
    warnings.filterwarnings('ignore')

    contractions_dict = {
    "ain't": "am not",
    "aren't": "are not",
    "can't": "cannot",
    "can't've": "cannot have",
    "'cause": "because",
    "could've": "could have",
    "couldn't": "could not",
    "couldn't've": "could not have",
    "didn't": "did not",
    "doesn't": "does not",
    "doesn’t": "does not",
    "don't": "do not",
    "don’t": "do not",
    "hadn't": "had not",
    "hadn't've": "had not have",
    "hasn't": "has not",
    "haven't": "have not",
    "he'd": "he had",
    "he'd've": "he would have",
    "he'll": "he will",
    "he'll've": "he will have",
    "he's": "he is",
    "how'd": "how did",
    "how'd'y": "how do you",
    "how'll": "how will",
    "how's": "how is",
    "i'd": "i would",
    "i'd've": "i would have",
    "i'll": "i will",
    "i'll've": "i will have",
    "i'm": "i am",
    "i've": "i have",
    "isn't": "is not",
    "it'd": "it would",
    "it'd've": "it would have",
    "it'll": "it will",
    "it'll've": "it will have",
    "it's": "it is",
    "let's": "let us",
    "ma'am": "madam",
    "mayn't": "may not",
    "might've": "might have",
    "mightn't": "might not",
    "mightn't've": "might not have",
    "must've": "must have",
    "mustn't": "must not",
    "mustn't've": "must not have",
    "needn't": "need not",
    "needn't've": "need not have",
    "o'clock": "of the clock",
    "oughtn't": "ought not",
    "oughtn't've": "ought not have",
    "shan't": "shall not",
    "sha'n't": "shall not",
    "shan't've": "shall not have",
    "she'd": "she would",
    "she'd've": "she would have",
    "she'll": "she will",
    "she'll've": "she will have",
    "she's": "she is",
    "should've": "should have",
    "shouldn't": "should not",
    "shouldn't've": "should not have",
    "so've": "so have",
    "so's": "so is",
    "that'd": "that would",
    "that'd've": "that would have",
    "that's": "that is",
    "there'd": "there would",
    "there'd've": "there would have",
    "there's": "there is",
    "they'd": "they would",
    "they'd've": "they would have",
    "they'll": "they will",
    "they'll've": "they will have",
    "they're": "they are",
    "they've": "they have",
    "to've": "to have",
    "wasn't": "was not",
    "we'd": "we would",
    "we'd've": "we would have",
    "we'll": "we will",
    "we'll've": "we will have",
    "we're": "we are",
    "we've": "we have",
    "weren't": "were not",
    "what'll": "what will",
    "what'll've": "what will have",
    "what're": "what are",
    "what's": "what is",
    "what've": "what have",
    "when's": "when is",
    "when've": "when have",
    "where'd": "where did",
    "where's": "where is",
    "where've": "where have",
    "who'll": "who will",
    "who'll've": "who will have",
    "who's": "who is",
    "who've": "who have",
    "why's": "why is",
    "why've": "why have",
    "will've": "will have",
    "won't": "will not",
    "won't've": "will not have",
    "would've": "would have",
    "wouldn't": "would not",
    "wouldn't've": "would not have",
    "y'all": "you all",
    "y’all": "you all",
    "y'all'd": "you all would",
    "y'all'd've": "you all would have",
    "y'all're": "you all are",
    "y'all've": "you all have",
    "you'd": "you would",
    "you'd've": "you would have",
    "you'll": "you will",
    "you'll've": "you will have",
    "you're": "you are",
    "you've": "you have",
    "ain’t": "am not",
    "aren’t": "are not",
    "can’t": "cannot",
    "can’t’ve": "cannot have",
    "’cause": "because",
    "could’ve": "could have",
    "couldn’t": "could not",
    "couldn’t’ve": "could not have",
    "didn’t": "did not",
    "doesn’t": "does not",
    "don’t": "do not",
    "don’t": "do not",
    "hadn’t": "had not",
    "hadn’t’ve": "had not have",
    "hasn’t": "has not",
    "haven’t": "have not",
    "he’d": "he had",
    "he’d’ve": "he would have",
    "he’ll": "he will",
    "he’ll’ve": "he will have",
    "he’s": "he is",
    "how’d": "how did",
    "how’d’y": "how do you",
    "how’ll": "how will",
    "how’s": "how is",
    "i’d": "i would",
    "i’d’ve": "i would have",
    "i’ll": "i will",
    "i’ll’ve": "i will have",
    "i’m": "i am",
    "i’ve": "i have",
    "isn’t": "is not",
    "it’d": "it would",
    "it’d’ve": "it would have",
    "it’ll": "it will",
    "it’ll’ve": "it will have",
    "it’s": "it is",
    "let’s": "let us",
    "ma’am": "madam",
    "mayn’t": "may not",
    "might’ve": "might have",
    "mightn’t": "might not",
    "mightn’t’ve": "might not have",
    "must’ve": "must have",
    "mustn’t": "must not",
    "mustn’t’ve": "must not have",
    "needn’t": "need not",
    "needn’t’ve": "need not have",
    "o’clock": "of the clock",
    "oughtn’t": "ought not",
    "oughtn’t’ve": "ought not have",
    "shan’t": "shall not",
    "sha’n’t": "shall not",
    "shan’t’ve": "shall not have",
    "she’d": "she would",
    "she’d’ve": "she would have",
    "she’ll": "she will",
    "she’ll’ve": "she will have",
    "she’s": "she is",
    "should’ve": "should have",
    "shouldn’t": "should not",
    "shouldn’t’ve": "should not have",
    "so’ve": "so have",
    "so’s": "so is",
    "that’d": "that would",
    "that’d’ve": "that would have",
    "that’s": "that is",
    "there’d": "there would",
    "there’d’ve": "there would have",
    "there’s": "there is",
    "they’d": "they would",
    "they’d’ve": "they would have",
    "they’ll": "they will",
    "they’ll’ve": "they will have",
    "they’re": "they are",
    "they’ve": "they have",
    "to’ve": "to have",
    "wasn’t": "was not",
    "we’d": "we would",
    "we’d’ve": "we would have",
    "we’ll": "we will",
    "we’ll’ve": "we will have",
    "we’re": "we are",
    "we’ve": "we have",
    "weren’t": "were not",
    "what’ll": "what will",
    "what’ll’ve": "what will have",
    "what’re": "what are",
    "what’s": "what is",
    "what’ve": "what have",
    "when’s": "when is",
    "when’ve": "when have",
    "where’d": "where did",
    "where’s": "where is",
    "where’ve": "where have",
    "who’ll": "who will",
    "who’ll’ve": "who will have",
    "who’s": "who is",
    "who’ve": "who have",
    "why’s": "why is",
    "why’ve": "why have",
    "will’ve": "will have",
    "won’t": "will not",
    "won’t’ve": "will not have",
    "would’ve": "would have",
    "wouldn’t": "would not",
    "wouldn’t’ve": "would not have",
    "y’all": "you all",
    "y’all": "you all",
    "y’all’d": "you all would",
    "y’all’d’ve": "you all would have",
    "y’all’re": "you all are",
    "y’all’ve": "you all have",
    "you’d": "you would",
    "you’d’ve": "you would have",
    "you’ll": "you will",
    "you’ll’ve": "you will have",
    "you’re": "you are",
    "you’re": "you are",
    "you’ve": "you have",
    }

    """Creating a regular expression that matches any of the contractions in the text.
    (can't|I'm|won't)
    Will generate an expression similar to this , to expand contractions in main text.
    """

    contractions_re = re.compile('(%s)' % '|'.join(contractions_dict.keys()))

    def cleanhtml(raw_html):
        cleanr = re.compile('<.*?>')
        # re.compile('') this compiles a regular expression to a regular expression object
        cleantext = re.sub(cleanr, '', raw_html)
        # sub method of re module to replace occurences of cleanr with empty string
        return cleantext

    def expand_contractions(s ,contractions_dict=contractions_dict):
        def replace(match):
            print(match)
            return contractions_dict[match.group(0)]
        # replace is a helper function that returs expanded form
        return contractions_re.sub(replace ,s)

    text=" i'm sleepy can't do more work"

    print(expand_contractions(text))

    def preprocessing(article):
        global article_sent

        # Converting to lowercase
        article = article.str.lower()

        # Removing the HTML
        article = article.apply(lambda x: cleanhtml(x))

        # Removing the email ids
        article = article.apply(lambda x: re.sub('\S+@\S+','', x))

        # Removing The URLS
        article = article.apply(lambda x: re.sub("((http\://|https\://|ftp\://)|(www.))+(([a-zA-Z0-9\.-]+\.[a-zA-Z]{2,4})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(/[a-zA-Z0-9%:/-_\?\.'~]*)?",'', x))

        # Removing the '\xa0'
        article = article.apply(lambda x: x.replace("\xa0", " "))
        # handles text formatting
        # non breaking spaces

        # Removing the contractions
        article = article.apply(lambda x: expand_contractions(x))

        # Stripping the possessives
        # John's -> John
        article = article.apply(lambda x: x.replace("'s", ''))
        article = article.apply(lambda x: x.replace('’s', ''))
        article = article.apply(lambda x: x.replace("\'s", ''))
        article = article.apply(lambda x: x.replace("\’s", ''))

        # Removing the Trailing and leading whitespace and double spaces
        article = article.apply(lambda x: re.sub(' +', ' ',x))

        # Copying the article for the sentence tokenization
        article_sent = article.copy()

        # Removing punctuations from the article
        article = article.apply(lambda x: ''.join(word for word in x if word not in punctuation))

        # Removing the Trailing and leading whitespace and double spaces again as removing punctuation might
        # Lead to a white space
        article = article.apply(lambda x: re.sub(' +', ' ',x))

        # Removing the Stopwords
        article = article.apply(lambda x: ' '.join(word for word in x.split() if word not in stop_words))

        return article

    def normalize(li_word):
    # li_word is a list of dictionaries
    # word:frequency
    # this function normalizes frequencies of words
        global normalized_freq
        normalized_freq = []
        for dictionary in li_word:
            max_frequency = max(dictionary.values())
            for word in dictionary.keys():
                dictionary[word] = dictionary[word]/max_frequency
            normalized_freq.append(dictionary)
        return normalized_freq

    def word_frequency(article_word):
        word_frequency = {}
        li_word = []
        for sentence in article_word:
            for word in word_tokenize(sentence):
                if word not in word_frequency.keys():
                    word_frequency[word] = 1
                else:
                    word_frequency[word] += 1
            li_word.append(word_frequency)
            word_frequency = {}
        normalize(li_word)
        return normalized_freq

    # Calculates frequencies of words in an article and returns normalized frequencies

    def sentence_score(li):
        global sentence_score_list
        sentence_score = {}
        sentence_score_list = []
        for list_, dictionary in zip(li, normalized_freq):
            for sent in list_:
                for word in word_tokenize(sent):
                    if word in dictionary.keys():
                        if sent not in sentence_score.keys():
                            sentence_score[sent] = dictionary[word]
                        else:
                            sentence_score[sent] += dictionary[word]
            sentence_score_list.append(sentence_score)
            sentence_score = {}
        return sentence_score_list

    # Basically hum ye dekh rhe hai ki konse sentence me most significant words hai
    # with respect to document
    # so that we are able to summarize tasks

    def sent_token(article_sent):
        sentence_list = []
        sent_token = []
        for sent in article_sent:
            token = sent_tokenize(sent)
            for sentence in token:
                token_2 = ''.join(word for word in sentence if word not in punctuation)
                token_2 = re.sub(' +', ' ',token_2)
                sent_token.append(token_2)
            sentence_list.append(sent_token)
            sent_token = []
        sentence_score(sentence_list)
        return sentence_score_list

    # Takes a list of sentences and the preprocesses it and calculates sentence score

    def summary(sentence_score_OwO):
        summary_list = []
        for summ in sentence_score_OwO:
            select_length = int(len(summ)*0.25)
            summary_ = nlargest(select_length, summ, key = summ.get)
            summary_list.append(".".join(summary_))
        return summary_list

    def make_series(art):
        global dataframe
        data_dict = {'article' : [art]}
        dataframe = pd.DataFrame(data_dict)['article']
        return dataframe

    def article_summarize(artefact):

        if type(artefact) != pd.Series:
            artefact = make_series(artefact)

        df = preprocessing(artefact)

        word_normalization = word_frequency(df)

        sentence_score_OwO = sent_token(article_sent)

        summarized_article = summary(sentence_score_OwO)

        return summarized_article

    def word_cloud(art):
        art_ = make_series(art)
        OwO = preprocessing(art_)
        wordcloud_ = WordCloud(height = 500, width = 1000, background_color = 'white').generate(art)
        plt.figure(figsize=(15, 10))
        plt.imshow(wordcloud_, interpolation='bilinear')
        plt.axis('off')
    # Generating the summaries for the first 100 articles
    # summaries = article_summarize(df['article'][0:100])

    text="'WASHINGTON  —   Congressional Republicans have a new fear when it comes to their    health care lawsuit against the Obama administration: They might win. The incoming Trump administration could choose to no longer defend the executive branch against the suit, which challenges the administration’s authority to spend billions of dollars on health insurance subsidies for   and   Americans, handing House Republicans a big victory on    issues. But a sudden loss of the disputed subsidies could conceivably cause the health care program to implode, leaving millions of people without access to health insurance before Republicans have prepared a replacement. That could lead to chaos in the insurance market and spur a political backlash just as Republicans gain full control of the government. To stave off that outcome, Republicans could find themselves in the awkward position of appropriating huge sums to temporarily prop up the Obama health care law, angering conservative voters who have been demanding an end to the law for years. In another twist, Donald J. Trump’s administration, worried about preserving executive branch prerogatives, could choose to fight its Republican allies in the House on some central questions in the dispute. Eager to avoid an ugly political pileup, Republicans on Capitol Hill and the Trump transition team are gaming out how to handle the lawsuit, which, after the election, has been put in limbo until at least late February by the United States Court of Appeals for the District of Columbia Circuit. They are not yet ready to divulge their strategy. “Given that this pending litigation involves the Obama administration and Congress, it would be inappropriate to comment,” said Phillip J. Blando, a spokesman for the Trump transition effort. “Upon taking office, the Trump administration will evaluate this case and all related aspects of the Affordable Care Act. ” In a potentially   decision in 2015, Judge Rosemary M. Collyer ruled that House Republicans had the standing to sue the executive branch over a spending dispute and that the Obama administration had been distributing the health insurance subsidies, in violation of the Constitution, without approval from Congress. The Justice Department, confident that Judge Collyer’s decision would be reversed, quickly appealed, and the subsidies have remained in place during the appeal. In successfully seeking a temporary halt in the proceedings after Mr. Trump won, House Republicans last month told the court that they “and the  ’s transition team currently are discussing potential options for resolution of this matter, to take effect after the  ’s inauguration on Jan. 20, 2017. ” The suspension of the case, House lawyers said, will “provide the   and his future administration time to consider whether to continue prosecuting or to otherwise resolve this appeal. ” Republican leadership officials in the House acknowledge the possibility of “cascading effects” if the   payments, which have totaled an estimated $13 billion, are suddenly stopped. Insurers that receive the subsidies in exchange for paying    costs such as deductibles and   for eligible consumers could race to drop coverage since they would be losing money. Over all, the loss of the subsidies could destabilize the entire program and cause a lack of confidence that leads other insurers to seek a quick exit as well. Anticipating that the Trump administration might not be inclined to mount a vigorous fight against the House Republicans given the  ’s dim view of the health care law, a team of lawyers this month sought to intervene in the case on behalf of two participants in the health care program. In their request, the lawyers predicted that a deal between House Republicans and the new administration to dismiss or settle the case “will produce devastating consequences for the individuals who receive these reductions, as well as for the nation’s health insurance and health care systems generally. ” No matter what happens, House Republicans say, they want to prevail on two overarching concepts: the congressional power of the purse, and the right of Congress to sue the executive branch if it violates the Constitution regarding that spending power. House Republicans contend that Congress never appropriated the money for the subsidies, as required by the Constitution. In the suit, which was initially championed by John A. Boehner, the House speaker at the time, and later in House committee reports, Republicans asserted that the administration, desperate for the funding, had required the Treasury Department to provide it despite widespread internal skepticism that the spending was proper. The White House said that the spending was a permanent part of the law passed in 2010, and that no annual appropriation was required  —   even though the administration initially sought one. Just as important to House Republicans, Judge Collyer found that Congress had the standing to sue the White House on this issue  —   a ruling that many legal experts said was flawed  —   and they want that precedent to be set to restore congressional leverage over the executive branch. But on spending power and standing, the Trump administration may come under pressure from advocates of presidential authority to fight the House no matter their shared views on health care, since those precedents could have broad repercussions. It is a complicated set of dynamics illustrating how a quick legal victory for the House in the Trump era might come with costs that Republicans never anticipated when they took on the Obama White House.'"

    print("length of text" , len(text))

    text_summarized=article_summarize(text)

    print(len(text_summarized[0]))
    print(text_summarized)
    text_summarized



    res = (article_summarize(text))
    return jsonify(res)

# POST route for spelling checking
@app.route('/grammar', methods=['POST'])
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

# GET route
@app.route('/recommend', methods=['POST'])
def get_data():
    data = request.json  # assuming the data is sent as JSON
    headline = data.get("headline", "")
    content = data.get("content", "")
    author = data.get("author", "")

    import numpy as np
    import pandas as pd

    import math
    import matplotlib.pyplot as plt
    import seaborn as sns
    import plotly.graph_objects as go


    from nltk.corpus import stopwords
    from nltk.tokenize import word_tokenize
    from nltk.stem import WordNetLemmatizer

    from sklearn.feature_extraction.text import CountVectorizer
    from sklearn.feature_extraction.text import TfidfVectorizer

    from sklearn.metrics.pairwise import cosine_similarity
    from sklearn.metrics import pairwise_distances

    news=pd.read_json('news.json' , lines=True)

    # news.info()

    # news.head()

    news=news[news['date']>=pd.Timestamp(2017,1,1)]

    # news.shape

    news=news[news['headline'].apply(lambda x: len(x.split())>5)]
    # removing headlines with less than 5 words: headlines may become empty due to removal of stop words

    news.drop_duplicates('headline' , inplace=True)

    # news.isna().sum()

    # news['category'].unique()

    # news['category'].value_counts().sort_values(ascending=False)

    news.index=range(news.shape[0])

    # news['day and month']=news['date'].dt.strftime("%a")+news['date'].dt.strftime("%b")

    news_temp=news.copy()

    import nltk
    nltk.download('stopwords')

    stop_words=set(stopwords.words('english'))
    # stop_words

    def preprocess_headline(headline):
        headline=''.join(e for e in headline if e.isalnum() or e.isspace()).lower()
        words=[word for word in headline.split() if word not in stop_words]
        return ' '.join(words)

    news_temp['headline']=news_temp['headline'].apply(preprocess_headline)

    import nltk
    nltk.download('punkt')

    import nltk
    nltk.download('wordnet')

    lem=WordNetLemmatizer()
    # Reduces words to their root form

    for i in range(len(news_temp["headline"])):
        string = ""
        for w in word_tokenize(news_temp["headline"][i]):
            string += lem.lemmatize(w,pos = "v") + " "
        news_temp.at[i, "headline"] = string.strip()
        if(i%1000==0):
            print(i)

    nltk.download('punkt')

    nltk.download('wordnet')

    # min_df=0

    tfidf_headline_vectorizer=TfidfVectorizer()
    tfidf_headline_features=tfidf_headline_vectorizer.fit_transform(news_temp['headline'])

    tfidf_desc_vectorizer=TfidfVectorizer()
    tfidf_desc_features=tfidf_desc_vectorizer.fit_transform(news_temp['short_description'])

    tfidf_author_vectorizer=TfidfVectorizer()
    tfidf_author_features=tfidf_author_vectorizer.fit_transform(news_temp['authors'])

    from scipy.sparse import hstack

    combined_features=hstack([
        tfidf_headline_features,
        tfidf_desc_features,
        tfidf_author_features,
        # tfidf_category_features
    ])

    pd.set_option('display.max_colwidth' , None)

    """Preparing functions for recommending top 5 articles"""

    import json

    def preprocess_text(headline):
        headline=''.join(e for e in headline if e.isalnum() or e.isspace()).lower()
        words=[word for word in headline.split() if word not in stop_words]
        return ' '.join(words)

    def vectorize_text(feature , vectorizer):
        tfidf_features=vectorizer.transform([feature])
        return tfidf_features

    def combine_features(new_headline , new_description , new_author):
        headline_features = vectorize_text(new_headline ,tfidf_headline_vectorizer)
        description_features = vectorize_text(new_description ,tfidf_desc_vectorizer)
        author_features = vectorize_text(new_author ,tfidf_author_vectorizer)
        combined = hstack([headline_features, description_features, author_features])
        return combined

    def recommend_similar_news(new_head , new_desc , new_author ,k):
        new_head=preprocess_text(new_head)
        new_desc=preprocess_text(new_desc)
        features_new_data=combine_features(new_head ,new_desc , new_author)
        couple_dist = pairwise_distances(combined_features,features_new_data)
        indices = np.argsort(couple_dist.ravel())[1:k+1]
        df = pd.DataFrame({'publish_date': news['date'][indices].dt.strftime('%Y-%m-%d').values,
                    'headline':news['headline'][indices].values,
                        'authors':news['authors'][indices].values,
                        'desc':news['short_description'][indices].values,
                        'category':news['category'][indices].values,
                        'Euclidean similarity with the queried article': couple_dist[indices].ravel()})
        result_dict = df.to_dict(orient='records')

            # Convert dictionary to JSON
        result_json = json.dumps(result_dict, indent=4)

        return result_json

    news.loc[12 ,'headline']

    news.loc[12 ,'short_description']
    demi = {"link":"https:\/\/www.huffpost.com\/entry\/covid-boosters-uptake-us_n_632d719ee4b087fae6feaac9","headline":"Over 4 Million Americans Roll Up Sleeves For Omicron-Targeted COVID Boosters","category":"U.S. NEWS","short_description":"Health experts said it is too early to predict whether demand would match up with the 171 million doses of the new boosters the U.S. ordered for the fall.","authors":"Carla K. Johnson, AP","date":1663891200000}
    # print(recommend_similar_news("Over 4 Million Americans Roll Up Sleeves For Omicron-Targeted COVID Boosters" ,"Health experts said it is too early to predict whether demand would match up with the 171 million doses of the new boosters the U.S. ordered for the fall.","Carla K. Johnson, AP",5))

    res = (recommend_similar_news(headline, content, author, 5))
    return jsonify(res)

if __name__ == '__main__':
    app.run(debug=True)
