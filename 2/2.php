<?php

add_action( 'wp_footer', function( $title ) {
	?>
	<p><?php echo 'This does not pass PHPCS. Missing sanitization.'; ?></p>
	<?php
} );

