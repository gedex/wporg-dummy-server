<?php

function _p1_wp_footer() {
	?>
	<p><?php echo esc_html( 'This pass PHPCS' ); ?></p>
	<?php
}

add_action( 'wp_footer', '_p1_wp_footer' );
