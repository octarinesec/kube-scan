import React from 'react';
import {Router, Route, browserHistory} from 'react-router';

import Risk from './Risk';

const Routes = props => {
    return (
        <Router history={browserHistory} {...props}>
            <Route path="/" component={Risk}/>
            <Route path="risk" component={Risk}/>
        </Router>
    );
};

export default Routes;
