import React, { useState } from "react";
import { createDocument, CreateDocumentPayload } from "../../../api/documentService";

type Props = {
    module_id: number;
}

const CreateDocumentCard = (props:Props) => {
  const [docTitle, setDocTitle] = useState("");
  const [docContent, setDocContent] = useState("");
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);    
  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      // Check if file size exceeds 10MB
      if (file.size > 10 * 1024 * 1024) {
        setErrorMessage("File size exceeds 10MB.");
        setSelectedFile(null);
        return;
      }

      // Check if the file is a PDF
      if (file.type !== "application/pdf") {
        setErrorMessage("Only PDF files are allowed.");
        setSelectedFile(null);
        return;
      }

      // Clear error message and set the file if all checks pass
      setErrorMessage(null);
      setSelectedFile(file);
    }
    else {
        setErrorMessage("No file selected.");
        setSelectedFile(null);
    }
  };


  const onCreateClick = async () => {
    if (!selectedFile) {
        setErrorMessage("Please select a valid file.");
        return;
      }
  
      const payload: CreateDocumentPayload = {
        doc_name: docTitle,
        doc_description: docContent
      };
    try {
        const response = await createDocument(props.module_id, selectedFile, payload);
        console.log(response);
      console.log("Doc Created");
    } catch (e) {
      if (e instanceof Error) {
        console.log(e.message);
      } else {
        console.log("An unknown error occurred");
      }
    }
  };
  return (
    <div
      className="
        h-fit w-full
        flex flex-col 
        bg-base-100
        text-base-content
        relative 
        text-left
        p-2
        text-wrap break-normal
        overflow-hidden
        text-ellipsis
        "
    >
      <h1 className="truncate h-fit mb-1">Create Document</h1>
      <span className="font-bold">Title:</span>
      <textarea
        className="text-wrap break-words h-20 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1          
          "
        onChange={(e) => setDocTitle(e.target.value)}
        value={docTitle}
      />
      <div
        className="flex flex-col"
        onSubmit={() => {
          console.log("submit");
        }}
      >
        <span className="font-bold">Description:</span>
        <textarea
          className="text-wrap break-words h-20 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1          
          "
          onChange={(e) => setDocContent(e.target.value)}
          value={docContent}
        />

        <label className="form-control w-full max-w my-2">
          <div className="label">
            <span className="label-text">Upload pdf</span>
            <span className="label-text-alt text-gray-400">Max file size: 10mb</span>
          </div>
          <input
            type="file"
            className="file-input file-input-bordered w-full max-w"
            onChange={handleFileChange}
          />
        </label>

        {errorMessage && <p className="text-red-500">{errorMessage}</p>}
        <button
            disabled={!selectedFile || errorMessage!==null}
          className="mb-2 p-6"
          // log the selected option
          onClick={(e) => {
            e.preventDefault();
            console.log("Update");
            onCreateClick();
          }}
        >
          Create Document
        </button>
        <button className="mb-1 p-4" type="submit">
          Cancel
        </button>
      </div>
    </div>
  );
};

export default CreateDocumentCard;
