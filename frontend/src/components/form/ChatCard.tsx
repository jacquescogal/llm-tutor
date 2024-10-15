import React, { useState, useEffect, useRef } from "react";
import { Chat, ENTITY_TYPE, ID_TYPE } from "../../types/chat";
import {
  createGeneration,
  createGenerationResponse,
} from "../../api/generationService";

type Props = {
  id: number;
  id_type: ID_TYPE;
};

const ChatCard = (props: Props) => {
  const [chatHistory, setChatHistory] = useState<Chat[]>([
    {
      content: "Hello there!",
      entity_type: ENTITY_TYPE.BOT,
    },
    {
      content: "Hello there!",
      entity_type: ENTITY_TYPE.USER,
    },
  ]);
  const [input, setInput] = useState<string>("");

  // Create a reference for the chat history div
  const chatHistoryRef = useRef<HTMLDivElement>(null);

  // Scroll to bottom when chat history updates
  useEffect(() => {
    if (chatHistoryRef.current) {
      chatHistoryRef.current.scrollTop = chatHistoryRef.current.scrollHeight;
    }
  }, [chatHistory]);

  const handleSubmission = async () => {
    console.log("Submit");
    const newChatHistory = [
      ...chatHistory,
      { content: input, entity_type: ENTITY_TYPE.USER },
    ];

    setInput("");
    // filtered chat is last 3 messages
    const filteredChat = newChatHistory.slice(-3);
    try {
      const response: createGenerationResponse = await createGeneration({
        id: props.id,
        id_type: props.id_type,
        chat_history: filteredChat,
      });
      newChatHistory.push({
        content: response.response,
        entity_type: ENTITY_TYPE.BOT,
      });
      setChatHistory(newChatHistory);
    } catch (e) {
      setChatHistory(newChatHistory);
      if (e instanceof Error) {
        console.log(e.message);
      } else {
        console.log("An unknown error occurred");
      }
    }
  };

  return (
    <>
      <h1 className="truncate h-fit mb-1 text-slate-600">Chat Context:{
      props.id_type === ID_TYPE.MODULE ? " Module" : props.id_type === ID_TYPE.SUBJECT ? " Subject" : " Document"
      } {String(props.id).padStart(5, "0")}
      
      </h1>
      <div className="h-96 overflow-y-scroll bg-gray-100 shadow-inset" ref={chatHistoryRef}>
        {chatHistory.map((chat, index) => {
          return (
            <>
              {chat.entity_type === ENTITY_TYPE.USER ? (
                <div className="chat chat-end" key={index}>
                  <div className="chat-bubble bg-blue-500">{chat.content}</div>
                </div>
              ) : (
                <div className="chat chat-start" key={index}>
                  <div className="chat-bubble bg-slate-600">{chat.content}</div>
                </div>
              )}
            </>
          );
        })}
      </div>
      <textarea
        className="text-wrap break-words h-20 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1  w-full   resize-none"
        onChange={(e) => setInput(e.target.value)}
        value={input}
      />
      <button
        className="py-4 w-full"
        onClick={(e) => {
          e.preventDefault();
          handleSubmission();
        }}
      >
        Submit
      </button>
    </>
  );
};

export default ChatCard;
