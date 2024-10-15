import React, { useRef } from "react";
import { useSelector } from "react-redux";
import { RootState } from "../../store/store";

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

  const loadingCount = useSelector((state: RootState) => state.isLoading.loadingCount);

  return (
    <>
      <button className={props.className} onClick={openModal}>
        {props.buttonName}
      </button>
      <dialog id="my_modal_1" className="modal  z-[5000]" ref={modalRef}>
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
        {loadingCount>0 && <div className="bg-gray-100 w-full h-full fixed top-0 left-0 z-[9999] opacity-50"/>}
      </dialog>
    </>
  );
};

export default ModalButton;
