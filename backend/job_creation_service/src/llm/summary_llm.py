from openai import AsyncOpenAI
from src.utils.configuration import Configuration
class SummaryLLM:
    def __init__(self):
        self.config = Configuration.get_instance()
        self.client = AsyncOpenAI(api_key=self.config.openai_api_key)
        self.model = self.config.openai_api_model

    async def summarise_document(self, document_text):
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
This is the document to summarise:
{document_text}"""}
            ],
            temperature=0
        )
        summary =  response.choices[0].message.content
        return summary.strip()
    async def create_memory_documents(self, document_text):
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
            "memory_card_header": "<related details title>",
            "memory_card_body": "<related details detailed body>"
        },
        ...
        ]
        '''

        # Call the OpenAI API
        response = await self.client.chat.completions.create(
            model=self.model,
            messages=[
                {"role": "system", "content": prompt},
                {"role": "user", "content": f"""
                 document_name: Marshall Wace Sustainable Investing and Stewardship Approach.pdf
                 This is the document to analyse and compartmentalize into related details:
        .\n\n{document_text}"""}
            ],
            temperature=0
        )

        # Extract the summary from the response
        summary =  response.choices[0].message.content
        
        return summary.strip()