from langchain_text_splitters.markdown import MarkdownTextSplitter
from langchain_text_splitters import RecursiveCharacterTextSplitter
from src.utils.chunk_util import ChunkUtil
import math
class SplitterUtil:
    @staticmethod
    def split_by_markdown(text):
        splitter = MarkdownTextSplitter()
        return splitter.split_text(text)
    
    @staticmethod
    def split_by_characters(text, token_size=512, chunk_overlap_ratio=0.25):
        splitter = RecursiveCharacterTextSplitter(
            chunk_size=token_size,
            chunk_overlap=math.ceil(token_size * chunk_overlap_ratio),
            length_function=ChunkUtil.ChunkTokenCount,
            is_separator_regex=False
        )
        return splitter.split_text(text)
    