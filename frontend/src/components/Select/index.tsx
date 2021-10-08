import React from "react";
import { useIntl, FormattedMessage } from "react-intl";
import Select, { components, Props } from "react-select";
import styled, { useTheme } from "styled-components";
import { IconSharpArrowdown } from "@subgame/react-icon-subgame";

export type TOption = {
  label: string;
  value: string | number;
};

interface MultiTagSelectProps extends Props {
  isError?: boolean;
  onChange?: any;
  options: any[];
  label?: string;
  labelPlaceholder?: boolean;
  value?: any;
  customHeight?: string;
  customWidth?: string;
}

const CustomSelect: React.FunctionComponent<MultiTagSelectProps> = ({
  label,
  options,
  defaultValue,
  customHeight,
  customWidth,
  menuIsOpen,
  isError,
  onFocus,
  onBlur,
  onChange,
  placeholder,
  labelPlaceholder,
  value,
}) => {
  const IconDropdown: React.FunctionComponent = () => (
    <IconSharpArrowdown color={theme.Secondary_Black} size={20} />
  );
  const { formatMessage } = useIntl();
  const theme = useTheme();
  const styles: any = {
    control: (provided: any, state: any) => {
      return {
        ...provided,
        width: customWidth ? customWidth : "100%",
        height: customHeight ? customHeight : "30px",
        minHeight: customHeight ? customHeight : "30px",
        padding: 0,
        borderRadius: "4px",
        background: "#FFF",
        borderColor: isError
          ? theme.Function_Error
          : `${theme.Secondary_Grey_3}!important;`,
        boxShadow: "none",
        boxSizing: "border-box",
        "*": {
          boxSizing: "border-box",
        },
      };
    },
    menuList: (provided: any) => {
      return {
        ...provided,
        padding: 0,
        background: theme.Pop_Up,
        zIndex: "500",
        borderRadius: "4px",
      };
    },
    container: (provided: any) => {
      return {
        ...provided,
        width: "100%",
      };
    },

    menu: () => {
      return {
        top: "calc(100% + 5px)",
        position: "absolute",
        left: 0,
        zIndex: "999",
        minWidth: "100%",
        boxShadow: "0px 0px 10px -2px rgba(0, 0, 0, 0.25);",
        padding: 0,
        borderRadius: "4px",
      };
    },

    option: (provided: any, state: any) => {
      return {
        ...provided,
        height: "40px",
        padding: " 0 15px",
        color: theme.Secondary_Black,
        backgroundColor: state.isSelected
          ? theme.Secondary_Grey_3
          : state.isFocused
          ? theme.Secondary_Grey_3
          : "#FFF",
        // "&:active": {
        //   backgroundColor: theme.primaryColor1,
        // },
        whiteSpace: "nowrap",
        display: "flex",
        justifyContent: "flex-start",
        alignItems: "center",
        fontWeight: 400,
        fontSize: "15px",
      };
    },

    singleValue: (provided: any) => {
      return {
        ...provided,
        color: theme.Secondary_Black,
        fontWeight: 400,
        fontSize: "15px",
      };
    },
    dropdownIndicator: (provided: any) => {
      return {
        ...provided,
        padding: "0 8px 0 0",
        cursor: "pointer",
        "i,i::before": {
          fontSize: "20px",
          color: theme.Secondary_Black,
        },
      };
    },
    menuPortal: (provided: any) => ({ ...provided, zIndex: 998 }),
  };

  const customComponents = {
    MultiValueRemove: () => null,
    DropdownIndicator: (props: any) => {
      return (
        <components.DropdownIndicator {...props}>
          <button type="button">
            <IconDropdown />
          </button>
        </components.DropdownIndicator>
      );
    },
    SingleValue: (props: any) => {
      return (
        <components.SingleValue {...props}>
          <FormattedMessage id={props?.data?.label || "All"} />
        </components.SingleValue>
      );
    },
    Option: (props: any) => {
      return (
        <SelectOptionMenuStyles>
          <components.Option {...props}>
            <label>
              <FormattedMessage id={props.label || "unknown"} />
            </label>
          </components.Option>
        </SelectOptionMenuStyles>
      );
    },
    CrossIcon: () => null,
    ClearIndicator: () => null,
    IndicatorSeparator: () => null,
  };

  const noOptionsMessage = (obj: any) => {
    return formatMessage({ id: "NoDataFound" });
  };

  return (
    <Body>
      {/* {label && <Label label={label} />}
      {!label && labelPlaceholder && (
        <Label label="label" style={{ opacity: "0" }} />
      )} */}
      <Select
        value={value}
        menuIsOpen={menuIsOpen}
        defaultValue={defaultValue}
        isSearchable={false}
        hideSelectedOptions={false}
        closeMenuOnSelect={true}
        menuPortalTarget={document.body}
        menuPosition={"absolute"}
        menuPlacement="auto"
        onFocus={onFocus}
        onBlur={onBlur}
        onChange={onChange}
        noOptionsMessage={noOptionsMessage}
        options={options}
        styles={styles}
        placeholder={formatMessage({
          id: placeholder?.toString() ?? "Select",
        })}
        components={customComponents}
      />
    </Body>
  );
};
const SelectOptionMenuStyles = styled.div`
  label {
    margin: 0px 10px 0 0px;
    font-weight: 400;
  }
`;

const Body = styled.div`
  display: flex;
  align-items: center;
  width: 100%;
  input {
    width: 100%;
  }
`;

export default CustomSelect;
