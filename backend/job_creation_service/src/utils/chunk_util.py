import tiktoken

class ChunkUtil:
    @staticmethod
    def ChunkTokenCount(chunk):
        tokenizer = tiktoken.encoding_for_model("gpt-4o-mini")
        tokens = tokenizer.encode(chunk)
        return len(tokens)