import React, {useEffect, useContext} from 'react';
import {Link} from 'react-router';
import App, {DataContext} from './App';
import './Risk.scss';
import OCTableCell from "./components/CommonTable/OCTableCell";
import K8sRisksTable from "./components/Risk/K8sRisksTable";
import { parseK8sRisksWorkloads } from "@octarine/ui-common";
import RefreshIcon from '@material-ui/icons/Refresh';
import CircularProgress from '@material-ui/core/CircularProgress';

function Risk() {
    return (
        <App>
            <DataContext.Consumer >
                {value => (
                    <div className="Risk oc-main-page">
                        <div className="Home-header">
                            <h2>K8S Risk Assessment</h2>
                            <button disabled={value.state.refreshing} onClick={value.onRefreshClick} className='refresh-state-btn'>
                              {value.state.refreshing ? (<CircularProgress size='16px' className='refreshIcon' />) : (<RefreshIcon className='refreshIcon' />)}
                              <span>{value.state.refreshing ? "Refreshing..." : "Refresh"}</span>
                            </button>
                        </div>
                        <K8sRisksTable
                            risks={value.state ? value.state['data']:null}
                            dispatch={value.dispatch}
                            currentSort={value.state ? value.state['sortField']:null}
                            ascending={value.state ? value.state['ascending']:null}
                        />
                    </div>
                )}

            </DataContext.Consumer>
        </App>
    );
}

export default Risk;
