<?php
$texto = "O dia da mentira é 01/04/2021\n";
$texto.= "O ultimo natal foi em 24/12/2020\n";

function prox_ano($padroes)
{
  return $padroes[1].($padroes[2]+1);
}
echo preg_replace_callback(
            "|(\d{2}/\d{2}/)(\d{4})|",
            "prox_ano",
            $texto);

?>