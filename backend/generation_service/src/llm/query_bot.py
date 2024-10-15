from langchain.prompts import SystemMessagePromptTemplate,HumanMessagePromptTemplate,ChatPromptTemplate
from langchain.chains import ConversationalRetrievalChain
from langchain.chat_models import ChatOpenAI
import os
from datetime import datetime
from src.utils.configuration import Configuration
from src.utils.retriever import QuestionRetriever
from src.protos.generation_pb2 import EntityType, ChatHistory
from typing import List


class QueryBot:
    instance=None
    def __init__(self):
        general_system_template = r""" 
        You will be asked questions by a student.
        You only reply in English.
        You should only use knowledge base provided below, date and chat history to answer the question.
        You are a tutor. If you do not know an answer, you should ask them to upload it to the database of documents first.
        Your purpose is to answer the question as best as you can to help the student understand.
        You should sound articulate, smart and capable in your answers. You should speak in the first person.
        You should limit your knowledge to the information below.
        Limit your answer to a 200 words.
        ----
        Date:
        {formatted_date}
        ----
        Knowledge Base:
        {context}
        ----
        Chat History:
        {chat_history}
        """
        general_user_template = "Question:```{question}```"
        messages = [
                    SystemMessagePromptTemplate.from_template(general_system_template),
                    HumanMessagePromptTemplate.from_template(general_user_template)
        ]
        self.qa_prompt = ChatPromptTemplate.from_messages( messages)
        self.ai_model = Configuration.get_instance().openai_api_model
        self.open_ai_key = Configuration.get_instance().openai_api_key
        
    @classmethod
    def get_instance(cls):
        if cls.instance is None:
            cls.instance = cls()
        return cls.instance
    
    def query(self, chat_history:List[ChatHistory], vector_id, id_type):
        print(chat_history)
        print(type(chat_history))
        memory=[]
        chat_pair={"user":None,"ai":None}
        try:
            qa=ConversationalRetrievalChain.from_llm(
                ChatOpenAI(openai_api_key=self.open_ai_key,temperature=0, model=self.ai_model),
                retriever=QuestionRetriever(k=5,vector_id=vector_id,id_type=id_type), 
                condense_question_llm = ChatOpenAI(openai_api_key=self.open_ai_key,temperature=0, model=self.ai_model),
                combine_docs_chain_kwargs={"prompt": self.qa_prompt})
        except Exception as e:
            print(e)
        try:
            # Get current date, month, and year
            current_date = datetime.now()
            current_day = current_date.day
            current_month_name = current_date.strftime("%B")
            current_year = current_date.year

            # Format the date as 'Day Month YYYY'
            formatted_date = f"{current_day} {current_month_name} {current_year}"
            last_message = ""
            for message in chat_history:
                if message.entity_type == EntityType.ENTITY_TYPE_USER:
                    chat_pair["user"]=message.content
                    last_message = message.content
                elif message.entity_type == EntityType.ENTITY_TYPE_BOT:
                    chat_pair["ai"]=message.content
                    if chat_pair["user"] is not None and chat_pair["ai"] is not None:
                        memory.append((chat_pair["user"],chat_pair["ai"]))
                        chat_pair={"user":None,"ai":None}
        except Exception as e:
            print(e)
        # print(memory)
        # print(last_message)
        try:
            result = qa({"question": last_message, "chat_history":memory, "formatted_date":formatted_date})
        except Exception as e:
            print(e)
        # print(result)
        return result["answer"]