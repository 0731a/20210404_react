import React from 'react';
import { BrowserRouter as router, Route } from 'react-router-dom';

import CardContainer from './ProductCards';
import Nav from './Navigation';
import { SignInModalWindow, BuyModalWindow } from '/.modalwindows';
import About from './About';
import Orders from './orders';



class App extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      user:{
        loggedin: false,
        name: ""
      }
    };
    this.showSignInModalWindow = this.showSignInModalWindow.bind(this);
    this.showBuyModalWindow = this.showBuyModalWindow.bind(this);
    this.toggleBuyModalWindow = this.toggleBuyModalWindow.bind(this);
    this.toggleSignInModalWindow = this.toggleSignInModalWindow(this);
  }

  showSignInModalWindow(){
    const state = this.state;
    const newState = Object.assign({},state,{showSignInModal:true});
    this.setState(newState);
  }
  showBuyModalWindow(){
    const state = this.state;
    const newState = Object.assign({},state,{showBuyModal:true, productid:id, price:price});
    this.setState(newState);
  }
  toggleSignInModalWindow(){
    const state = this.state;
    const newState = Object.assign({},state,{showSignInModal: !state.showSignInModal });
    this.setState(newState);
  }
  toggleBuyModalWindow(){
    const state = this.state;
    const newState = Object.assign({},state,{showBuyModal:!state.showBuyModal});
    this.setState(newState);
  }
    
  render(){
    return (
      <div >
        <Router>
          <div>
            <Nav user={this.state.user} showModalWindow={this.showSignInModalWindow} />
            <div className="container pt-4 mt-4">
              <Router exact path="/" render={()=> <CardContainer location="cards.json" showBuyModal={this.showBuyModalWindow}/>} />
              <Router path="/promos" render={()=> <CardContainer location="promos.json" promo={true}  showBuyModal={this.showBuyModalWindow}/>} />
              { this.state.user.loggedin? <Router path="/myorders" render={()=> <Orders/>} />: null }
              <Router path="/about" Component={About}/>
            </div>
            <SignInModalWindow showModal={this.state.showSignInModal} toggle={this.toggleSignInModalWindow} />
            <BuyModalWindow showModal={this.state.showBuyModal} toggle={this.toggleBuyModalWindow} proudctid={this.state.productid} price={this.state.price} />
          </div>
        </Router>
      </div>
    );
  }
}

export default App;
