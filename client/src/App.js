import React, { useReducer, useEffect } from "react";
import "./fonts/fonts.scss";
import "./index.scss";
import "./App.scss";
import Toolbar from "./components/Toolbar/Toolbar";
import runtimeConfig from "./config";
import BottomBar from "./components/BottomBar/BottomBar";
import { get as lookupGet } from "lodash";
import { RiskBreakdownPopup, parseK8sRisksWorkloads } from "@octarine/ui-common";
import "@octarine/ui-common/dist/main.css";

let cNames = "OCApp side-menu-on";

export const DataContext = React.createContext(null);

let secondarySortField = "risk.riskScore";

let sortData = (sortField, ascending) => {
  return (a, b) => {
    function compare(field) {
      let aVal = lookupGet(a, field),
        bVal = lookupGet(b, field);
      if (aVal < bVal) {
        return ascending ? -1 : 1;
      } else if (aVal === bVal) {
        return 0;
      } else {
        return ascending ? 1 : -1;
      }
    }

    let primaryResult = compare(sortField);
    if (primaryResult !== 0) {
      return primaryResult;
    } else {
      if (secondarySortField === sortField) {
        return primaryResult;
      } else {
        return compare(secondarySortField);
      }
    }
  };
};

function reducer(state, action) {
  switch (action.type) {
    case "set":
      let newState = {
        data: action.data.data,
        sortField: state.sortField,
        ascending: state.ascending,
        popupOn: state.popupOn
      };
      if (newState.data) {
        newState.data.sort(sortData(state.sortField, state.ascending));
      }

      return newState;
    case "sort":
      if (state.sortField === action.sortField) {
        state.ascending = !state.ascending;
      }
      state.data.sort(sortData(action.sortField, state.ascending));
      return {
        data: state.data,
        sortField: action.sortField,
        ascending: state.ascending
      };
    case "popup":
      return {
        data: state.data,
        sortField: state.sortField,
        ascending: state.ascending,
        popupOn: true,
        popupData: action.riskData
      };
    case "closePopup":
      return {
        data: state.data,
        sortField: state.sortField,
        ascending: state.ascending,
        popupOn: false
      };
    default:
      throw new Error();
  }
}

const initialState = {
  data: null,
  sortField: "risk.riskScore",
  ascending: false,
  popupOn: false
};

function App(props) {
  const [state, dispatch] = useReducer(reducer, initialState);

  useEffect(() => {
    async function fetchData() {
      const result = await fetch("/api/risks");
      const { data } = await result.json();

      dispatch({
        type: "set",
        data: { data: parseK8sRisksWorkloads(data) }
      });
    }
    fetchData();
  }, []);

  function closePopup() {
    dispatch({
      type: "closePopup"
    });
  }
  return (
    <div className={cNames}>
      <Toolbar contactLink={runtimeConfig.contactLink} />
      <div className="app-main-row">
        <DataContext.Provider value={{ state, dispatch }}>
          <div className="current-page-wrapper">{props.children}</div>
        </DataContext.Provider>
      </div>
      <BottomBar websiteLink={runtimeConfig.websiteLink} />
      {state.popupOn ? (
        <RiskBreakdownPopup
          workload={state.popupData}
          onClose={closePopup}
        ></RiskBreakdownPopup>
      ) : null}
    </div>
  );
}

export default App;
