import React from "react";
import OCTableHeader from "../../components/CommonTable/OCTableHeader";
import OCTableHeaderCell from "../../components/CommonTable/OCTableHeaderCell";
import OCTableHoverRow from "../../components/CommonTable/OCTableHoverRow";
import OCTableCell from "../../components/CommonTable/OCTableCell";
import OCVirtualTable from "../../components/CommonTable/OCVirtualTable";
import './K8sRisksTable.scss';
import cn from "classnames";

function K8sRisksTable({ risks, openPopup, sortFunc, currentSort, ascending }) {
    function _renderHeader() {
        const tableId = "risksTable";
        return (
            <OCTableHeader>
                {/*
                 // @ts-ignore */}
                <OCTableHeaderCell
                    isSortable={true}
                    fieldName={'risk'}
                    tableID={tableId}
                    currentSort={currentSort}
                    ascending={ascending}
                    sortField={'risk.riskScore'}
                    sortFunc={sortFunc}>Risk</OCTableHeaderCell>
                {/*
                 // @ts-ignore */}
                <OCTableHeaderCell
                    isSortable={true}
                    fieldName={'resourceName'}
                    tableID={tableId}
                    currentSort={currentSort}
                    ascending={ascending}
                    sortField={'name'}
                    sortFunc={sortFunc}>Name</OCTableHeaderCell>
                {/*
                 // @ts-ignore */}
                <OCTableHeaderCell
                    isSortable={true}
                    fieldName={'resourceKind'}
                    tableID={tableId}
                    currentSort={currentSort}
                    ascending={ascending}
                    sortField={'kind'}
                    sortFunc={sortFunc}>Kind</OCTableHeaderCell>
                {/*
                 // @ts-ignore */}
                <OCTableHeaderCell
                    isSortable={true}
                    fieldName={'resourceNamespace'}
                    tableID={tableId}
                    currentSort={currentSort}
                    ascending={ascending}
                    sortField={'namespace'}
                    sortFunc={sortFunc}>Namespace</OCTableHeaderCell>
            </OCTableHeader>
        );
    }

    function _renderRow (_, item) {
        let roundScore = Math.round(item.risk.riskScore);

        return (
            <OCTableHoverRow onClick={() => openPopup(item)}>
                <OCTableCell classNames={[item.risk.riskCategory.toString().toLowerCase()]} fieldName={'risk'}><div>{roundScore}</div></OCTableCell>
                <OCTableCell fieldName={'resourceName'}>{item.name}</OCTableCell>
                <OCTableCell fieldName={'resourceKind'}>{item.kind}</OCTableCell>
                <OCTableCell fieldName={'resourceNamespace'}>{item.namespace}</OCTableCell>
            </OCTableHoverRow>
        );
    }

    const cNames = cn(['K8sRisksTable', 'oc-main-page']);

    return (
        <div className={cNames}>
            <OCVirtualTable
                overscanRowCount={500}
                renderRow={_renderRow}
                renderHeader={_renderHeader}
                classNames={['CommonTable']}
                items={risks}
            />
        </div>
    );
}

export default K8sRisksTable;