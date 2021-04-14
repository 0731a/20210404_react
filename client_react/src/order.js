import React from 'react';

function Order(props){
    return(
        <div className="col-12">
            <div className="card text-center">
                <div className="cardheader"><h5>{props.productname}</h5></div>
                <div className="card-body">
                    <div className="row">
                        <div className="mx-auto col-6">
                            <img src={props.img} alt={props.imgalt} className="img-thumbnail float-left" />
                        </div>
                        <div className="col-6">
                            <p className="card-text">{props.desc}</p>
                            <div className="mb-4">
                                Price: <strong>{props.price}</strong>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="card-footer text-muted">
                    Purchase {props.days} days ago
                </div>
            </div>
            <div className="mt-3" />
        </div>
    );
}

export default class OrderContainer extends React.Component{
    constructor(props){
        super(props);
        this.state = {
            orders: []
        };
    }

    render(){
        return(
            <div className="row mt-5">
                {this.state.orders.map(order => { <Order key={order.id} {...order} />})}
            </div>
        );
    }
}