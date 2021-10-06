// Cheditor5 不提供TS定義檔案，在這邊手動加入避免ts檢查
// Empty typings for the editor used in the app to satisfy the TS compiler in the strict mode:
declare module '@ckeditor/ckeditor5-build-classic' {
  // or other CKEditor 5 build.
  const ClassicEditorBuild: any

  export = ClassicEditorBuild
}

declare module '@ckeditor/ckeditor5-react' {
  // or other CKEditor 5 build.
  const ClassicEditorBuild: any

  export = ClassicEditorBuild
}

declare module '@ckeditor/ckeditor5-custom-build' {
  // or other CKEditor 5 build.
  const ClassicEditorBuild: any

  export = ClassicEditorBuild
}

declare module '@ckeditor/ckeditor5-custom-build/build/ckeditor' {
  // or other CKEditor 5 build.
  const ClassicEditorBuild: any

  export = ClassicEditorBuild
}
