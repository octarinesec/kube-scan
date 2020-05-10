import React, { useEffect, useState } from "react";
import "./fonts/fonts.scss";
import "./index.scss";
import "./App.scss";
import Toolbar from "./components/Toolbar/Toolbar";
import runtimeConfig from "./config";
import BottomBar from "./components/BottomBar/BottomBar";
import { get as lookupGet } from "lodash";
import { RiskBreakdownPopup, parseK8sRisksWorkloads } from "@octarine/ui-common";
import "@octarine/ui-common/dist/main.css";

const getRefreshStatusIntervalSeconds = 3
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

function App(props) {
  const [refreshing, setRefreshing] = useState(false);
  const [risksData, setRisksData] = useState(null);
  const [sortField, setSortField] = useState("risk.riskScore");
  const [ascending, setAscending] = useState(false);
  const [selectedShowSystemNamespaces, setSelectedShowSystemNamespaces] = useState(false)
  const [popupData, setPopupData] = useState(null);

  async function fetchData() {
    const result = await fetch("/api/risks");
    const {data} = await result.json();
    setRisksData(data)
  }

  async function updateRefreshingStatus(lastFetch) {
    const result = await fetch("/api/refreshing_status");
    const {refreshing, lastRefresh} = await result.json();
    setRefreshing(refreshing)
    if (lastRefresh > lastFetch) {
      lastFetch = Date.now()
      await fetchData()
    }
    await new Promise(resolve => setTimeout(resolve, getRefreshStatusIntervalSeconds * 1000));
    await updateRefreshingStatus(lastFetch)
  }

  useEffect(() => {
    fetchData().then()
    updateRefreshingStatus(Date.now()).then()
  }, []);

  async function refreshState() {
    setRefreshing(true)
    const result = await fetch("/api/refresh", {method: 'post'});
    await result.json();
  }

  function openPopup(newPopupData) {
    setPopupData(newPopupData)
  }

  function closePopup() {
    setPopupData(null)
  }

  let risks = risksData ? [...risksData] : null
  if (risks) {
    if (!selectedShowSystemNamespaces) {
      risks = risks.filter(r => !r.isSystemWorkload)
    }
    risks = parseK8sRisksWorkloads(risks)
    risks.sort(sortData(sortField, ascending))
  }

  function sortFunc(newSortField) {
    if (sortField === newSortField) {
      setAscending(!ascending);
    }
    setSortField(newSortField)
  }

  return (
    <div className={ cNames }>
      <Toolbar contactLink={ runtimeConfig.contactLink } />
      <div className="app-main-row">
        <DataContext.Provider value={ {risks, sortField, ascending, sortFunc, openPopup, refreshState, refreshing, selectedShowSystemNamespaces, setSelectedShowSystemNamespaces} }>
          <div className="current-page-wrapper">{ props.children }</div>
        </DataContext.Provider>
      </div>
      <BottomBar websiteLink={ runtimeConfig.websiteLink } />
      { popupData ? (
        <RiskBreakdownPopup
          workload={ popupData }
          onClose={ closePopup }
        />
      ) : null }
    </div>
  );
}

export default App;
