import React from 'react';
import { Modal, ModalHeader, ModalBody} from 'reactstrap'; 

export function BuyModalWindow (props){
    
    return(
        <modal id="buy" tabIndex="-1" role="dialog" isOpen={props.showModal} toggle={props.togle}>
            <div role="document">
                <ModalHeader toggle={props.toggle} className="bg-success text-white">
                    Buy Item
                </ModalHeader>
                <ModalBody>
                    {/*신용카드 결제 폼*/}
                </ModalBody>
            </div>
        </modal>
    );
    
}