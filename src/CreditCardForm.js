import React from 'react';
import { injectStripe, StripeProvider, Elements, CardElement } from 'react-stripe-elements';

// 신용카드 처리 과정 상태
const INITIALSTATE = "INITIAL", SUCCESSTATE = "COMPLETE", FAILEDSTATE = "FAILED";


class CreditCardForm extends React.Component{
    constructor(props){ 
        super(props);
        this.state = { status: INITIALSTATE };    
    }

    /*
        1. 입력된 신용카드를 나타내는 토큰을 발급한다.
        2. 백엔드 서버로 토큰을 발급한다.
        3. 결제 성공이나 실패 여부에 맞는 화면을 렌더링
    */
    async handleSubmit(event){
        event.preventDefault();
        console.log("Handle submit called, with name: "+this.state.value);
        // Stripe API를 통해 토큰 발급
        let { token } = await this.props.stripe.createToken({ name: this.state.value });
        if(token == null){
            console.log("invalid token");
            this.setState({ status: FAILEDSTATE });
            return;
        }
        let response = await fetch("/charge",{  // 해당 URL로 POST 요청 송신
            method: "POST",
            headers: { "Content-Type": "text/plain"},
            body: JSON.stringify({
                token: token.id,
                operation: this.props.operation,
            })
        });
        console.log(response.ok);
        if(response.ok) {
            console.log("Purchase Complete!");
            this.setState({ status: SUCCESSTATE });
        }
    }

    renderCreditCardInformation() {
        const style={
            base: {
                'fontsize': '20px',
                'color': '#495057',
                'fontFamily': 'apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "HelveticaNeue" Arial, sans-serif'
            }
        };

        const usersavedcard = <div>
            <div className="form-row text-center">
                <button type="button" className="btn btn-outline-success text-center mx-auto">Use saved card?</button>
            </div>
            <hr />
        </div>;

        const remembercardcheck = <div className="form-row form-check textcenter">
            <input className="form-check-input" type="checkbox" value="" id="remembercardcheck" />
            <lable className="form-check-label" htmlFor="remembercardcheck">
                Remember Card?
            </lable>
        </div>;

        return (
            <div>
                {usersavedcard}
                <h5 className="mb-4">Payment Info</h5>
                <form onSubmit={this.handleSubmit}>
                    <div className="form-row">
                        <div className="col-lg-12 form-group">
                            <label htmlFor="cc-name">Name On Card:</label>
                            <input id="cc-name" name='cc-name' className="form-control" placeholder='Name on Card' onChange={this.handleInputChange} type='text' />
                        </div>
                    </div>
                    <div className="form-row">
                        <div className="col-lg-12 form-group">
                            <label htmlFor="card">Card Information:</label>
                            <CardElement id="card" className="form-control" sytle={style} />
                        </div>
                    </div>
                    {remembercardcheck}
                    <hr className="mb-4" />
                    <button type="submit" className="btn btn-success btnlarge">{this.props.operation}</button> 
                </form>
            </div>
        );
    }
    renderSuccess() {
        return(
            <div>
                <h5 className="mb-4 text-success">Request Successfull....</h5>
                <button type="submit" className="btn btn-success btn-large" onClick={() => { this.props.toggle(); }}>Ok</button>
            </div>
        );
    }
    renderFailure() {
        return(
            <div>
                <h5 className="mb-4 text-danger"> Creadit card information invalid, try again or exit</h5>
                {this.renderCreditCardInformation()}
            </div>
        );
    }
    render(){
        let body = null;
        switch (this.state.status){
            case SUCCESSTATE:
                body = this.renderSuccess();
                break;
            case FAILEDSTATE:
                body = this.renderFailure();
                break;
            default:
                body = this.renderCreditCardInformation();
        }

        return(
            <div>
                {body}
            </div>
        )
    };
}
