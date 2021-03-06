export type Styles = {
  Pink_1: "#F72677";
  Pink_2: "#EB027D";
  text_1: "#98959E";
  gray_1: "#e7eaf3";
  gray_2: "#B9B9B9";
  gray_3: "#f3f5f9";
  gray_4: "#c0c4cc";

  blue_1: "#004DE1";
  //

  Primary_2: "#AD065E";

  Primary_3: "#EE72B4";

  Primary_4: "#FFDEEF";
  Primary_Gradient: "linear-gradient(309.48deg, #EB027D 0%, #FFDEEF 101.58%)";
  Secondary_Black: "#171717";
  Secondary_Grey_1: "#8B8B8B";
  Secondary_Grey_2: "#B9B9B9";

  Secondary_Grey_3: "#E8E8E8";
  Bg: "#F8F9FA";
  Dark_1: "#171717";
  Dark_2: " #1E2232";
  Dark_3: "#35394A";
  Dark_4: "#626572";
  Function_Error: "#AD065E";
  Function_complete: "#1162FF";
  Function_Green: "#25A68E";

  Board: "0px 0px 10px -2px rgba(0, 0, 0, 0.15)";
  Pop_Up: " 0px 10px 25px -5px rgba(0, 0, 0, 0.26)";

  fontStyles: {
    "60": `font-size:60px;font-weight:300;line-height:120%;letter-spacing:0.1em;`;
    "48": `font-size:48px;font-weight:300;line-height:120%;letter-spacing:0.1em;`;
    "36": `font-size:36px;font-weight:300;line-height:120%;letter-spacing:0.1em;`;
    "24": `font-size:24px;font-weight:400;line-height:120%;letter-spacing:0.05em;`;
    "18": `font-size:18px;font-weight:600;line-height:160%;letter-spacing:0.05em;`;
    "16": `font-size:16px;font-weight:400;line-height:160%;letter-spacing:0.05em;`;
    "14": `font-size:14px;font-weight:600;line-height:21px;letter-spacing:0.05em;`;
    "13": `font-size:13px;font-weight:600;line-height:18px;letter-spacing:0.1em;`;
    "12": `font-size:12px;font-weight:400;line-height:18px;letter-spacing:0.1em;`;
  };
};

declare module "styled-components" {
  export interface DefaultTheme extends Styles {}
}
