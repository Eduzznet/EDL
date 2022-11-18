<?php
$numDado = mt_rand(1, 6);
$numExt = "";

  switch($numDado) 
  {
  case 1:
    $numExt = "<i>um</i>";
    break;
  case 2:
    $numExt = "<i>dois</i>";
    break;
  case 3:
    $numExt = "<i>tres</i>";
    break;
  case 4:
    $numExt = "<i>quatro</i>";
    break;
  case 5:
    $numExt = "<i>cinco</i>";
    break;
  case 6:
    $numExt = "<i>seis</i>";
    break;
  default:
    $numExt = "ERRO";
  }
  echo "<h1>Toda vez que atualizar a página o dado será jogado.</h1>";
  echo 'O dado deu '.$numExt.'.';

?>