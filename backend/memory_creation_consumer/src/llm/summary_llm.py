from openai import AsyncOpenAI
from src.utils.configuration import Configuration
import json
class SummaryLLM:
    def __init__(self):
        self.config = Configuration.get_instance()
        
        self.client = AsyncOpenAI(api_key=self.config.openai_api_key)
        self.model = self.config.openai_api_model

    async def summarise_document(self, documet_name, document_text):
        prompt = f"""
        Do not fabricate any information, only use the information provided in the document.
        Extract pertinent details from the document that a learner should know and would be useful in a knowledge base to ace an exam on the topic.
        These will be used inform the learner about the topic so all details should be included in a structured manner.
        Pertinent details include:
        - terms and definitions
        - important concepts
        - examples
        - diagrams
        - formulas
        - equations
        - tables
        These will be used inform the learner about the topic so all pertinent details should be included.
        Limit the output to 4,000 words.
        """

        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": prompt},
                {"role": "user", "content": f"""
                document_name: {documet_name}
                This is the document to summarise:
                {document_text}"""}
            ],
            temperature=0
        )
        summary =  response.choices[0].message.content
        return summary.strip()
    
    async def create_memory_documents(self, documet_name, document_text):
        # Prepare the prompt for the model
        prompt = '''
        Do not fabricate any information, only use the information provided in the document.
        Extract pertinent details from the document that a learner should know and would be useful in a knowledge base to ace an exam on the topic.
        These will be used inform the learner about the topic so all details should be included in a structured manner.
        Pertinent details include:
        - terms and definitions
        - important concepts
        - examples
        - diagrams
        - formulas
        - equations
        - tables
        Use only the provided document and limit output to 4,000 words.
        Compartmentalise into related granular details.
        Output template:
        [
        {
            "memory_title": "<related details title>",
            "memory_content": "<related details detailed body>"
        },
        ...
        ]
        '''

        template = """
        [
    {
        "memory_title": "<related details title>",
        "memory_content": "<related details detailed body>"
    },
    ...
    ]
    """

        # Call the OpenAI API
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": prompt},
                {"role": "user", "content": f"""
                 document_name: {documet_name}
                 This is the document to analyse and compartmentalize into related details:
        .\n\n{document_text}"""}
            ],
            temperature=0
        )

        # Extract the summary from the response
        summary =  response.choices[0].message.content
        
        return await self.to_json(summary.strip(), template)

    async def create_mcq_question(self, documet_name, document_text, question_count):
        # Prepare the prompt for the model
        prompt = f'''
    Do not fabricate any information, only use the information provided in the document.
    Create Multi Answer Multiple Choice Questions from the document that a learner should review and would be useful to learn from and study to ace an exam on the topic.
    These will be used for the learner to review.
    Pertinent questions include but are not limited to the following:
    - terms and definitions
    - important concepts
    - examples
    - diagrams
    - formulas
    - equations

    Number of questions to create: {question_count}
    Of 4 choices, at exactly 1 should be correct.

    Output template:
    [
    {{
        "question": "<question>",
        "choices": [
            {{
                "choice": "<choice 1>",
                "is_correct": true
            }},
            {{
                "choice": "<choice 2>",
                "is_correct": false
            }},
            {{
                "choice": "<choice 3>",
                "is_correct": false
            }},
            {{
                "choice": "<choice 4>",
                "is_correct": false
            }}
        ]
    }},
    ...
    ]
    '''
        template = """
        [
    {
    "question": "<question>",
    "choices": [
        {
            "choice": "<choice 1>",
            "is_correct": true
        },
        {
            "choice": "<choice 2>",
            "is_correct": false
        },
        {
            "choice": "<choice 3>",
            "is_correct": true
        },
        {
            "choice": "<choice 4>",
            "is_correct": false
        },
    ]
    }, 
    ...
    ]
    """

        # Call the OpenAI API
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": prompt},
                {"role": "user", "content": f"""
                 document_name: {documet_name}
                 This is the document to create questions out of:
        .\n\n{document_text}"""}
            ],
            temperature=0
        )

        # Extract the summary from the response
        summary =  response.choices[0].message.content
        
        return await self.to_json(summary.strip(), template)
    
    async def create_multi_answer_question(self, documet_name, document_text, question_count):
        prompt = f'''
    Do not fabricate any information, only use the information provided in the document.
    Create Multi Answer Multiple Choice Questions from the document that a learner should review and would be useful to learn from and study to ace an exam on the topic.
    These will be used for the learner to review.
    Pertinent questions include but are not limited to the following:
    - terms and definitions
    - important concepts
    - examples
    - diagrams
    - formulas
    - equations

    Number of questions to create: {question_count}
    Of 4 choices, at least 2 should be correct.

    Output template:
    [
    {{
        "question": "<question>",
        "choices": [
            {{
                "choice": "<choice 1>",
                "is_correct": true
            }},
            {{
                "choice": "<choice 2>",
                "is_correct": false
            }},
            {{
                "choice": "<choice 3>",
                "is_correct": true
            }},
            {{
                "choice": "<choice 4>",
                "is_correct": false
            }}
        ]
    }},
    ...
    ]
    '''

        template = """
        [
    {
    "question": "<question>",
    "choices": [
        {
            "choice": "<choice 1>",
            "is_correct": true
        },
        {
            "choice": "<choice 2>",
            "is_correct": false
        },
        {
            "choice": "<choice 3>",
            "is_correct": true
        },
        {
            "choice": "<choice 4>",
            "is_correct": false
        },
    ]
    }, 
    ...
    ]
    """


        # Call the OpenAI API
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": prompt},
                {"role": "user", "content": f"""
                 document_name: {documet_name}
                 This is the document to create questions out of:
        .\n\n{document_text}"""}
            ],
            temperature=0
        )

        # Extract the summary from the response
        summary =  response.choices[0].message.content
        
        return await self.to_json(summary.strip(), template)
    
    async def create_open_answer_question(self, documet_name, document_text, question_count):
        # Prepare the prompt for the model
        prompt = f'''
        Do not fabricate any information, only use the information provided in the document.
        Create flash card question and answer cards from the document that a learner should review and would be useful to learn from and study to ace an exam on the topic.
        These will be used for the learner to review.
        Pertinent questions include but is not limited to the following:
        - terms and definitions
        - important concepts
        - examples
        - diagrams
        - formulas
        - equations
        number of questions to create: {question_count}
        Output template:
        [
        {{
            "question": "<question>",
            "answer": "<answer>"
        }},
        ...
        ]
        '''
        template = """
        [
    {
    "question": "<question>",
    "answer": "<answer>"
    },
    ...
    ]
    """

        # Call the OpenAI API
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": prompt},
                {"role": "user", "content": f"""
                 document_name: {documet_name}
                 This is the document to create questions out of:
        .\n\n{document_text}"""}
            ],
            temperature=0
        )

        # Extract the summary from the response
        summary =  response.choices[0].message.content
        
        return await self.to_json(summary.strip(), template)
    
    async def to_json(self, text, template_text):
        print(f"Text: {text}")
        attempt_count = 3
        try:
            json_result =  json.loads(text)
            return json_result
        except Exception as e:
            while attempt_count > 0:
                try:
                    attempt_count -= 1
                    text = await self.generate_clean_json(text, template_text)
                    json_result =  json.loads(text)
                    return json_result
                except Exception as e:
                    print(f"Error parsing the summary: {text}")
                    continue
        return None
        
    async def generate_clean_json(self, dirty_json, template_text):
        prompt = f"""
        Clean the following JSON and make it valid:
        {dirty_json}
        """
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": prompt},
                {"role": "user", "content": f"This is the template to clean the JSON: {template_text}"}
            ],
            temperature=0
        )
        return response.choices[0].message.content.strip()