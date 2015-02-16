<?php

add_action( 'wp_footer', function( $title ) {
	?>
	<p><?php echo esc_html( 'This pass PHPCS' ); ?></p>
	<?php
} );

