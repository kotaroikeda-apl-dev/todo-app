import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';

// カスタムテーマの作成
const theme = createTheme({
  palette: {
    primary: {
      main: '#1976d2', // プライマリカラーの設定
    },
    secondary: {
      main: '#dc004e', // セカンダリカラーの設定
    },
  },
  typography: {
    fontFamily: 'Noto Sans JP, Arial', // 日本語フォントの設定
  },
});

ReactDOM.render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <CssBaseline /> {/* ベースラインのCSSを適用 */}
      <App />
    </ThemeProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

