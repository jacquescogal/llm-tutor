import React, { useRef } from "react";

type Props = {
  buttonName: string;
  children: React.ReactNode;
  className?: string;
};

const ModalButton = (props: Props) => {
  const modalRef = useRef<HTMLDialogElement>(null);

  const openModal = () => {
    if (modalRef.current) {
      modalRef.current.showModal();
    }
  };


  return (
    <>
      <button className={props.className} onClick={openModal}>
        {props.buttonName}
      </button>
      <dialog id="my_modal_1" className="modal" ref={modalRef}>
        <div className="modal-box">
          <div className="modal-action">
            <form className="w-full" method="dialog">
            {props.children}
              <button className="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">
                ✕
              </button>
              {/* <button className="btn">Close</button> */}
            </form>
          </div>
        </div>
      </dialog>
    </>
  );
};

export default ModalButton;
