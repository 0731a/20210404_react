import React from 'react';
import { Modal, ModalHeader, ModalBody} from 'reactstrap'; 
import CreditCardInformation from './CreditCard';


export function BuyModalWindow (props){
    
    return(
        <modal id="buy" tabIndex="-1" role="dialog" isOpen={props.showModal} toggle={props.toggle}>
            <div role="document">
                <ModalHeader toggle={props.toggle} className="bg-success text-white">
                    Buy Item
                </ModalHeader>
                <ModalBody>
                    <CreditCardInformation show={true} operation="Charge" toggle={props.toggle} />
                </ModalBody>
            </div>
        </modal>
    );
    
}

export class SignInModalWindow extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            showRegisterationForm : false // 해당 값에 따라 SignInform이나 RegisterationForm 컴포넌트를 모달 윈도우에 추가한다.
        };
        this.handlerNewUser = this.handlerNewUser.bind(this);
    }

    handlerNewUser(){
        this.setState({
            showRegisterationForm: true
        });
    }

    render(){
        let modalBody = <SignInForm handlerNewUser={this.handlerNewUser} />
        if( this.state.showRegisterationForm === true ){
            modalBody = <Registeration />
        }

        return (
            <modal id="register" tabIndex="-1" role="dialog" isOpen={this.props.showModal} toggle={this.props.toggle}>
                <div role="document">
                    <ModalHeader toggle={this.props.toggle} className="bg-success text-white">
                        Sign in
                        {
                            /* 
                                <button className="close">
                                    <span aria-hidden="true>&times;</span>
                                </button>
                            */
                        }
                    </ModalHeader>
                    <ModalBody>
                        {modalBody}
                    </ModalBody>
                </div>
            </modal>

        );
    }
}

class SignInForm extends React.Component{
    constructor(props){
        super(props);
        this.state = {
            errormessage:''
        };
        this.handleChange = this.handlerChange.bind(this); // 사용자가 데이터 입력시 호출 함수 
        this.handleSubmit = this.handlerSubmit.bind(this); // 폼을 제출하면 호출되는 함수 
    }

    handlerChange(event){
        const name = event.target.name;
        const value = event.target.value;
        this.setState({
            [name]: value
        });
    }

    handlerSubmit(event){
        event.preventDefault();
        console.log(JSON.stringify(this.state));
    }

    render(){
        let message = null;
        if( this.state.errormessage.length != 0 ){
            message = <h5 className="mb-4 textdanger">{this.state.errormessage}</h5>
        }
        return(
            <div>
                {message}
                <form onSubmit={this.handleSubmit}>
                    <h5 className="mb-4">Basic Info</h5>
                    <div className="form-group">
                        <label htmlFor="email">Email:</label>
                        <input name="email" type="email" className="form-control" id="email" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                            <label htmlFor="pass">Password:</label>
                            <input name="password" type="password" className="form-control" id="pass" onChange={this.handleChange} />
                    </div>
                    <div className="form-row text-center">
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btnsuccess btn-large">SignIn</button>
                        </div>
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btn-link text-info" onClick={()=>this.props.handlerNewUser()}> New User? Register</button>
                        </div>
                    </div>
                </form>
            </div>
        );
    }
}


class Registeration extends React.Component{
    constructor(props){
        super(props);
        this.handleChange = this.handlerChange.bind(this); // 사용자가 데이터 입력시 호출 함수 
        this.handleSubmit = this.handlerSubmit.bind(this); // 폼을 제출하면 호출되는 함수 
        this.state = {
            errormessage:''
        };
    }

    handlerChange(event){
        const name = event.target.name;
        const value = event.target.value;
        this.setState({
            [name]: value
        });
    }

    handlerSubmit(event){
        event.preventDefault();
        console.log(JSON.stringify(this.state));
    }

    render(){
        let message = null;
        if( this.state.errormessage.length != 0 ){
            message = <h5 className="mb-4 textdanger">{this.state.errormessage}</h5>
        }
        return(
            <div>
                {message}
                <form onSubmit={this.handleSubmit}>
                    <h5 className="mb-4">Registeration</h5>
                    <div className="form-group">
                        <label htmlFor="username">User Name:</label>
                        <input name="username" type="text" className="form-control" id="username" placeholder='John Doe' onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                        <label htmlFor="email">Email:</label>
                        <input name="email" type="email" className="form-control" id="email" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                            <label htmlFor="pass">Password:</label>
                            <input name="pass1" type="password" className="form-control" id="pass1" onChange={this.handleChange} />
                    </div>
                    <div className="form-group">
                            <label htmlFor="pass">Confirm Password:</label>
                            <input name="pass2" type="password" className="form-control" id="pass2" onChange={this.handleChange} />
                    </div>
                    <div className="form-row text-center">
                        <div className="col-12 mt-2">
                            <button type="submit" className="btn btnsuccess btn-large">Register</button>
                        </div>
                    </div>
                </form>
            </div>
        );
    }
}