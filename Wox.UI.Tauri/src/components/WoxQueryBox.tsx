import React, { useImperativeHandle } from "react"
import styled from "styled-components"
import { WoxTauriHelper } from "../utils/WoxTauriHelper.ts"
import { Theme } from "../entity/Theme.typings"
import { WoxThemeHelper } from "../utils/WoxThemeHelper.ts"

export type WoxQueryBoxRefHandler = {
  changeQuery: (query: string) => void
  selectAll: () => void
  focus: () => void
  getQuery: () => string
}

export type WoxQueryBoxProps = {
  defaultValue?: string
  onQueryChange: (query: string) => void
  onFocus?: () => void
  onClick?: () => void
}

export default React.forwardRef((_props: WoxQueryBoxProps, ref: React.Ref<WoxQueryBoxRefHandler>) => {
  const queryBoxRef = React.createRef<HTMLInputElement>()

  const selectInputText = () => {
    queryBoxRef.current?.select()
  }

  useImperativeHandle(ref, () => ({
    changeQuery: (query: string) => {
      if (queryBoxRef.current) {
        queryBoxRef.current!.value = query
        _props.onQueryChange(query)
      }
    },
    selectAll: () => {
      selectInputText()
    },
    focus: () => {
      queryBoxRef.current?.focus()
    },
    getQuery: () => {
      return queryBoxRef.current?.value ?? ""
    }
  }))

  return <Style theme={WoxThemeHelper.getInstance().getTheme()} className="wox-query-box">
    <input ref={queryBoxRef}
           title={"Query Wox"}
           className={"mousetrap"}
           type="text"
           aria-label="Wox"
           autoComplete="off"
           autoCorrect="off"
           autoFocus={true}
           autoCapitalize="off"
           defaultValue={_props.defaultValue}
           onFocus={() => {
             _props.onFocus?.()
           }}
           onClick={() => {
             _props.onClick?.()
           }}
           onChange={(e) => {
             _props.onQueryChange(e.target.value)
           }}

    />
    <div className={"dragging-container"}
         onMouseMoveCapture={(event) => {
           WoxTauriHelper.getInstance().startDragging().then(_ => {
             queryBoxRef.current?.focus()
           })
           event.preventDefault()
           event.stopPropagation()
         }}>&nbsp;</div>
  </Style>
})

const Style = styled.div<{ theme: Theme }>`
  position: relative;
  width: 100%;
  overflow: hidden;

  input {
    height: 60px;
    line-height: 60px;
    width: 100%;
    font-size: 24px;
    outline: none;
    padding-left: 10px;
    border: 0;
    cursor: auto;
    color: ${props => props.theme.QueryBoxFontColor};
    background-color: ${props => props.theme.QueryBoxBackgroundColor};
    border-radius: ${props => props.theme.QueryBoxBorderRadius}px;
    display: inline-block;
    box-sizing: border-box;
  }

  .dragging-container {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    background-color: transparent;
    width: 120px;
    z-index: 999;
  }
`