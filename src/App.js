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
  }
    
  render(){
    return (
      <div >
        <Router>
          <div>
            <Nav user={this.state.user} />
            <div className="container pt-4 mt-4">
              <Router exact path="/" render={()=> <CardContainer location="cards.json"/>} />
              <Router path="/promos" render={()=> <CardContainer location="promos.json" promo={true}/>} />
              { this.state.user.loggedin? <Router path="/myorders" render={()=> <Orders/>} />: null }
              <Router path="/about" Component={About}/>
            </div>
          </div>
        </Router>
      </div>
    );
  }
}

export default App;
